class ArgoCDAppManager {
    String argocdServer
    String argocdToken
    String project
    String helmChart
    String appName
    String helmValuesFile
    String destServer
    String destNamespace
    String helmRepoUrl
    String revision
    String argocdAppFile
    String appParams

    ArgoCDAppManager(String argocdServer, String argocdToken, String project, String appName, String destServer, String destNamespace, String helmChart = '', String helmValuesFile = '', String helmRepoUrl = '', String revision = '', String argocdAppFile = '', String appParams = '') {
        this.argocdServer = argocdServer
        this.argocdToken = argocdToken
        this.project = project
        this.appName = appName
        this.destServer = destServer
        this.destNamespace = destNamespace
        this.helmChart = helmChart
        this.helmValuesFile = helmValuesFile
        this.helmRepoUrl = helmRepoUrl
        this.revision = revision
        this.argocdAppFile = argocdAppFile
        this.appParams = appParams
    }

    void executeCommand(String command) {
        Process process = command.execute()
        process.waitFor()
        String output = process.in.text
        String error = process.err.text
        if (process.exitValue() != 0) {
            throw new RuntimeException("Error executing command: $error")
        }
        println(output)
    }

    void createWithHelm() {
        String helmSetParameters = helmValuesFile ? "--values ${helmValuesFile.split(',').join(' --values ')}" : ""
        String command = "argocd app create $appName --grpc-web --upsert --project $project --server $argocdServer --auth-token $argocdToken --repo $helmRepoUrl --revision $revision ${helmSetParameters} --helm-pass-credentials --dest-server $destServer --dest-namespace $destNamespace --path $helmChart"
        executeCommand(command)
    }

    void createFromfile() {
        executeCommand("argocd app create --grpc-web --server $argocdServer --auth-token $argocdToken --file $argocdAppFile")
    }

    void sync() {
        executeCommand("argocd app sync --grpc-web $appName --server $argocdServer --auth-token $argocdToken")
    }

    void waitForStatus() {
        executeCommand("argocd app wait --grpc-web --timeout 600 $appName --server $argocdServer --auth-token $argocdToken")
    }

    void delete() {
        executeCommand("argocd app delete --grpc-web $appName --server $argocdServer --auth-token $argocdToken")
    }

    void setAppParameters() {
        String params = appParams.split('\n').collect { pair ->
            def (key, value) = pair.split(':').collect { it.trim() }
            "--parameter ${key}=${value}"
        }.join(' ')
        executeCommand("argocd app set $appName --grpc-web --server $argocdServer --auth-token $argocdToken ${params}")
    }
}

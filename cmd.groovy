import groovy.json.JsonSlurper

class ArgoCDHelper {
    def argocdServer
    def argocdToken
    def argocdProject
    def helmChart
    def appName
    def helmValuesFile
    def destServer
    def destNamespace
    def helmRepoUrl
    def revision

    ArgoCDHelper(String argocdServer, String argocdToken, String argocdProject, String helmChart, String appName, String helmValuesFile, String destServer, String destNamespace, String helmRepoUrl, String revision) {
        this.argocdServer = argocdServer
        this.argocdToken = argocdToken
        this.argocdProject = argocdProject
        this.helmChart = helmChart
        this.appName = appName
        this.helmValuesFile = helmValuesFile
        this.destServer = destServer
        this.destNamespace = destNamespace
        this.helmRepoUrl = helmRepoUrl
        this.revision = revision
    }

    def deployHelmChart() {
        println "Deploying helm chart $helmChart to $destServer/$destNamespace using ArgoCD $argocdServer/$argocdProject"

        println "--server $argocdServer"
        println "--auth-token $argocdToken"
        println "--project $argocdProject"
        println "--repo $helmRepoUrl"
        println "--revision $revision"
        println "--dest-server $destServer"
        println "--dest-namespace $destNamespace"
        println "--path $helmChart"
        println "--name $appName"

        // Prepare the helm values file
        println "Preparing helm values file"
        def helmSetParameters = "--values"
        def files = helmValuesFile.split(',')
        for (file in files) {
            helmSetParameters += " ${file}"
        }

        println helmSetParameters

        // Create the application in ArgoCD
        def cmd = "argocd app create $appName --gprc-web --upsert --project $argocdProject --server $argocdServer --auth-token $argocdToken --repo $helmRepoUrl --revision $revision ${helmSetParameters} --helm-pass-credentials --dest-server $destServer --dest-namespace $destNamespace --path $helmChart"
        def process = cmd.execute()
        process.waitFor()

        if (process.exitValue() != 0) {
            throw new RuntimeException("Failed to deploy helm chart $helmChart to $destServer/$destNamespace using ArgoCD $argocdServer/$argocdProject")
        }
    }

    def setAppParameters(String appParams) {
        // Parse the helm values file map string into key=value pairs
        def params = ""
        def jsonSlurper = new JsonSlurper()
        def map = jsonSlurper.parseText(appParams)
        map.each { key, value ->
            params += " --parameter ${key}=${value}"
        }

        // Set the parameters for the application in ArgoCD
        def cmd = "argocd app set $appName --gprc-web --server $argocdServer --auth-token $argocdToken ${params}"
        def process = cmd.execute()
        process.waitFor()

        if (process.exitValue() != 0) {
            throw new RuntimeException("Failed to set parameters for application $appName in ArgoCD $argocdServer/$argocdProject")
        }
    }
}

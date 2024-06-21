class ArgoCDHelmDeployer {
    String argoServer
    String argoToken
    String helmChartPath
    String helmReleaseName
    String helmNamespace

    ArgoCDHelmDeployer(String argoServer, String argoToken, String helmChartPath, String helmReleaseName, String helmNamespace) {
        this.argoServer = argoServer
        this.argoToken = argoToken
        this.helmChartPath = helmChartPath
        this.helmReleaseName = helmReleaseName
        this.helmNamespace = helmNamespace
    }

    def deploy() {
        def command = "argocd app create ${helmReleaseName} --repo ${helmChartPath} --path . --dest-namespace ${helmNamespace} --dest-server ${argoServer}"
        def process = command.execute()
        process.waitFor()

        if(process.exitValue() == 0) {
            println "Application created successfully"
        } else {
            println "Error in creating application: ${process.err.text}"
        }
    }
}

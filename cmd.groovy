class ArgoCDHelmDeployer {
    String argoServer
    String argoToken
    String helmChartPath
    String helmValuesPath
    String helmReleaseName
    String helmNamespace
    String helmVersion
    Map<String, Object> helmValues
    String syncPolicy

    ArgoCDHelmDeployer(String argoServer, String argoToken, String helmChartPath, String helmValuesPath, String helmReleaseName, String helmNamespace, String helmVersion, Map<String, Object> helmValues, String syncPolicy) {
        this.argoServer = argoServer
        this.argoToken = argoToken
        this.helmChartPath = helmChartPath
        this.helmValuesPath = helmValuesPath
        this.helmReleaseName = helmReleaseName
        this.helmNamespace = helmNamespace
        this.helmVersion = helmVersion
        this.helmValues = helmValues
        this.syncPolicy = syncPolicy
    }

    def deploy() {
        def command = "argocd app create ${helmReleaseName} --repo ${helmChartPath} --path . --values ${helmValuesPath} --dest-namespace ${helmNamespace} --dest-server ${argoServer} --helm-set-string helmVersion=${helmVersion} --sync-policy ${syncPolicy} --auth-token ${argoToken}"
        helmValues.each { key, value ->
            command += " --helm-set-string ${key}=${value}"
        }
        def process = command.execute()
        process.waitFor()

        if(process.exitValue() == 0) {
            println "Application created successfully"
        } else {
            println "Error in creating application: ${process.err.text}"
        }
    }
}

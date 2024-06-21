import io.argoproj.argocd.v1alpha1.api.AppsApi
import io.argoproj.argocd.v1alpha1.model.App
import io.argoproj.argocd.v1alpha1.model.AppDestinationServer
import io.argoproj.argocd.v1alpha1.model.AppProjectSpec
import io.argoproj.argocd.v1alpha1.model.AppSourcePath
import io.argoproj.argocd.v1alpha1.model.AppSpec
import io.argoproj.argocd.v1alpha1.model.AppSyncPolicyAutomated
import io.argoproj.argocd.v1alpha1.model.RepoCreds
import io.kubernetes.client.util.Config

class ArgoCdHelmDeployer {
    private String namespace
    private String releaseName
    private String chartPath
    private String valuesFilePath
    private String argocdServer
    private String argocdToken
    private String argocdProject
    private String repoUrl

    ArgoCdHelmDeployer(String namespace, String releaseName, String chartPath, String valuesFilePath, String argocdServer, String argocdToken, String argocdProject, String repoUrl) {
        this.namespace = namespace
        this.releaseName = releaseName
        this.chartPath = chartPath
        this.valuesFilePath = valuesFilePath
        this.argocdServer = argocdServer
        this.argocdToken = argocdToken
        this.argocdProject = argocdProject
        this.repoUrl = repoUrl
    }

    def deploy() {
        // Configurer le client Argo CD
        ApiClient apiClient = Config.defaultClient()
        apiClient.setBasePath(argocdServer)
        apiClient.setAccessToken(argocdToken)
        AppsApi appsApi = new AppsApi(apiClient)

        // Créer l'objet App pour Argo CD
        App app = new App()
        AppSpec appSpec = new AppSpec()
        AppProjectSpec projectSpec = new AppProjectSpec()
        projectSpec.setName(argocdProject)
        appSpec.setProject(projectSpec)

        AppSourcePath sourcePath = new AppSourcePath()
        sourcePath.setChart(chartPath)
        sourcePath.setTargetRevision(valuesFilePath)
        sourcePath.setRepoURL(repoUrl)
        appSpec.setSource(sourcePath)

        AppDestinationServer destinationServer = new AppDestinationServer()
        destinationServer.setServer("https://kubernetes.default.svc")
        destinationServer.setNamespace(namespace)
        appSpec.setDestination(destinationServer)

        AppSyncPolicyAutomated syncPolicy = new AppSyncPolicyAutomated()
        syncPolicy.setPrune(true)
        syncPolicy.setSelfHeal(true)
        appSpec.setSyncPolicy(syncPolicy)

        app.setSpec(appSpec)
        app.setMetadata(new io.kubernetes.client.openapi.models.V1ObjectMeta().name(releaseName))

        // Créer ou mettre à jour l'application Argo CD
        try {
            appsApi.createApp(app)
            println("L'application Argo CD a été créée avec succès.")
        } catch (ApiException e) {
            if (e.getCode() == 409) {
                // L'application existe déjà, la mettre à jour
                appsApi.replaceApp(releaseName, app)
                println("L'application Argo CD a été mise à jour avec succès.")
            } else {
                throw e
            }
        }
    }
}

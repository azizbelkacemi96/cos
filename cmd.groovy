import io.argoproj.argocd.client.ApiClient
import io.argoproj.argocd.client.api.AppsApi
import io.argoproj.argocd.client.model.AppProject
import io.argoproj.argocd.client.model.HelmChart
import io.argoproj.argocd.client.model.RepoCreds
import io.argoproj.argocd.client.model.V1alpha1Application
import io.argoproj.argocd.client.model.V1alpha1ApplicationSpec
import io.argoproj.argocd.client.model.V1alpha1ApplicationSyncPolicy

class ArgoCdHelmDeployer {
    private ApiClient apiClient
    private AppsApi appsApi

    ArgoCdHelmDeployer(String argocdServer, String argocdToken) {
        this.apiClient = new ApiClient()
        this.apiClient.setBasePath(argocdServer)
        this.apiClient.setAccessToken(argocdToken)
        this.appsApi = new AppsApi(apiClient)
    }

    void deploy(String namespace, String releaseName, String chartPath, String valuesFilePath, String argocdProject, String repoUrl) {
        // Créer l'objet AppProject pour Argo CD
        AppProject appProject = new AppProject()
        appProject.setName(argocdProject)

        // Créer l'objet HelmChart pour Argo CD
        HelmChart helmChart = new HelmChart()
        helmChart.setChart(chartPath)
        helmChart.setValues(valuesFilePath)

        // Créer l'objet V1alpha1Application pour Argo CD
        V1alpha1Application app = new V1alpha1Application()
        V1alpha1ApplicationSpec appSpec = new V1alpha1ApplicationSpec()
        appSpec.setProject(appProject)
        appSpec.setSource(new io.argoproj.argocd.client.model.V1alpha1ApplicationSource().helm(helmChart).repoURL(repoUrl))
        appSpec.setDestination(new io.argoproj.argocd.client.model.V1alpha1ApplicationDestination().server("https://kubernetes.default.svc").namespace(namespace))
        appSpec.setSyncPolicy(new V1alpha1ApplicationSyncPolicy().automated(new io.argoproj.argocd.client.model.V1alpha1ApplicationSyncPolicyAutomated().prune(true)))
        app.setMetadata(new io.kubernetes.client.openapi.models.V1ObjectMeta().name(releaseName))
        app.setSpec(appSpec)

        // Créer ou mettre à jour l'application Argo CD
        try {
            appsApi.createApplication(app)
            println("L'application Argo CD a été créée avec succès.")
        } catch (ApiException e) {
            if (e.getCode() == 409) {
                // L'application existe déjà, la mettre à jour
                appsApi.updateApplication(releaseName, app)
                println("L'application Argo CD a été mise à jour avec succès.")
            } else {
                throw e
            }
        }
    }
}

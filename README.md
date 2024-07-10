argocdCreateAppWithHelm: creates a new ArgoCD application using a Helm chart. This method takes the following parameters:
ARGOCD_SERVER: the URL of the ArgoCD server
HELM_REPO_URL: the URL of the Helm repository containing the application chart
HELM_CHART_PATH: the path to the application chart in the Helm repository
APP_NAME: the name of the application to create in ArgoCD
ARGOCD_PROJECT_NAME: the name of the ArgoCD project in which the application will be created
REVISION: the revision of the application chart to deploy
DEST_SERVER: the URL of the target Kubernetes cluster
DEST_NAMESPACE: the target Kubernetes namespace where the application will be deployed
HELM_VALUES_FILES: the path to the Helm values files used to configure the application to be deployed
ARGOCD_APP_FILE: the path to the YAML file containing the definition of the ArgoCD application
APP_PARAMS: additional parameters to pass to the application deployment
argocdAppDelete: deletes an existing ArgoCD application. This method takes the following parameters:
ARGOCD_SERVER: the URL of the ArgoCD server
APP_NAME: the name of the application to delete in ArgoCD
ARGOCD_PROJECT_NAME: the name of the ArgoCD project in which the application is created
argocdAppSync: synchronizes an existing ArgoCD application with the target Kubernetes cluster. This method takes the following parameters:
ARGOCD_SERVER: the URL of the ArgoCD server
APP_NAME: the name of the application to synchronize in ArgoCD
ARGOCD_PROJECT_NAME: the name of the ArgoCD project in which the application is created
argocdAppStatus: retrieves the status of an existing ArgoCD application. This method takes the following parameters:
ARGOCD_SERVER: the URL of the ArgoCD server
APP_NAME: the name of the application whose status is to be retrieved in ArgoCD
ARGOCD_PROJECT_NAME: the name of the ArgoCD project in which the application is created
argocdCreateAppFromFile: creates a new ArgoCD application using a YAML file. This method takes the following parameters:
ARGOCD_SERVER: the URL of the ArgoCD server
APP_NAME: the name of the application to create in ArgoCD
ARGOCD_PROJECT_NAME: the name of the ArgoCD project in which the application will be created
ARGOCD_APP_FILE: the path to the YAML file containing the definition of the ArgoCD application
argocdSetAppParameters: sets the parameters of an existing ArgoCD application. This method takes the following parameters:
ARGOCD_SERVER: the URL of the ArgoCD server
APP_NAME: the name of the application whose parameters are to be set in ArgoCD
ARGOCD_PROJECT_NAME: the name of the ArgoCD project in which the application is created
APP_PARAMS: the parameters to set for the ArgoCD application

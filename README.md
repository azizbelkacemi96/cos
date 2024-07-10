# Jenkinsfile for Image Promotion and Application Deployment with ArgoCD

This Jenkinsfile utilizes a shared library called `sharedlib-cd` to provide two main functionalities:

1. **Image Promotion**
2. **Application Deployment with ArgoCD**

## Prerequisites

Before using this Jenkinsfile, make sure you have installed the following plugins in your Jenkins instance:

- **Credentials Binding Plugin**
- **Pipeline: Groovy**
- **Shared Libraries**

## Image Promotion

Image promotion uses the `promote` function from the `sharedlib-cd` shared library. This function takes the following parameters:

- `ARTIFACTORY_REGISTRY`: URL of the Artifactory registry
- `IMAGE_SOURCE`: Name of the source image to promote
- `IBM_CONTAINER_REGISTRY`: URL of the IBM Container registry
- `IBMSID_PATH`: Path to the IBM SID file
- `IMAGE_TARGET`: Name of the target image to promote
- `VAULT_NAMESPACE`: Vault namespace for authentication
- `VAULT_AUTH_TOKEN`: Vault authentication token
- `<ECOSYSTEM>_promote_ci_project_id`: CI project ID to promote. Replace `<ECOSYSTEM>` with the appropriate value for your ecosystem. See the **Finding the CI Project ID** section for details.
- `GITLAB_PROJECT_BRANCH`: GitLab project branch to promote
- `PROMOTE_CI_TRIGGER_TOKEN`: Trigger token for the promotion CI pipeline
- `VAULT_ADDR`: Vault address for authentication
- `PROMOTE_CI_API_TOKEN`: API token for the promotion CI pipeline

### Finding the CI Project ID

To find the CI project ID to promote, follow these steps:

1. Go to the CI project page in GitLab.
2. Click the "Settings" button in the left menu.
3. Click the "CI/CD" button in the left menu.
4. Copy the value of "Project ID" into the "Promote CI Project ID" field in your Jenkins pipeline.

## Application Deployment with ArgoCD

Application deployment with ArgoCD uses several functions from the `sharedlib-cd` shared library. All these functions take the following parameters:

- `ARGOCD_SERVER`: ArgoCD server URL
- `HELM_REPO_URL`: Helm repository URL for deployment
- `HELM_CHART_PATH`: Path to the Helm chart for deployment
- `APP_NAME`: Name of the application to deploy
- `ARGOCD_PROJECT_NAME`: ArgoCD project name for deployment
- `REVISION`: Helm chart revision for deployment
- `DEST_SERVER`: Destination server name for deployment
- `DEST_NAMESPACE`: Destination namespace for deployment
- `HELM_VALUES_FILES`: Helm values files for deployment
- `ARGOCD_APP_FILE`: ArgoCD application configuration file to use

Feel free to customize this content based on your specific configuration. If you have any other questions, I'm here to assist! 😊

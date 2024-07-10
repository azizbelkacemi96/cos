# Jenkinsfile pour la promotion d'image et le déploiement d'application avec ArgoCD

Ce Jenkinsfile utilise une shared library appelée `sharedlib-cd` pour fournir deux fonctionnalités principales :

1. **Promotion d'image**
2. **Déploiement d'application avec ArgoCD**

## Prérequis

Avant d'utiliser ce Jenkinsfile, assurez-vous d'avoir installé les plugins suivants dans votre instance Jenkins :

- **Credentials Binding Plugin**
- **Pipeline: Groovy**
- **Shared Libraries**

## Promotion d'image

La promotion d'image utilise la fonction `promote` de la shared library `sharedlib-cd`. Cette fonction prend les paramètres suivants :

- `ARTIFACTORY_REGISTRY` : l'URL du registre Artifactory
- `IMAGE_SOURCE` : le nom de l'image source à promouvoir
- `IBM_CONTAINER_REGISTRY` : l'URL du registre IBM Container
- `IBMSID_PATH` : le chemin d'accès au fichier IBM SID
- `IMAGE_TARGET` : le nom de l'image cible à promouvoir
- `VAULT_NAMESPACE` : l'espace de noms Vault à utiliser pour l'authentification
- `VAULT_AUTH_TOKEN` : le jeton d'authentification Vault à utiliser pour l'authentification
- `<ECOSYSTEM>_promote_ci_project_id` : l'ID du projet CI à promouvoir. `<ECOSYSTEM>` doit être remplacé par la valeur appropriée pour votre écosystème. Voir la section **Trouver l'ID du projet CI** pour plus de détails.
- `GITLAB_PROJECT_BRANCH` : la branche du projet GitLab à promouvoir
- `PROMOTE_CI_TRIGGER_TOKEN` : le jeton de déclenchement pour le pipeline CI de promotion
- `VAULT_ADDR` : l'adresse Vault à utiliser pour l'authentification
- `PROMOTE_CI_API_TOKEN` : le jeton d'API pour le pipeline CI de promotion

### Trouver l'ID du projet CI

Pour trouver l'ID du projet CI à promouvoir, suivez ces étapes :

1. Accédez à la page du projet CI dans GitLab.
2. Cliquez sur le bouton "Settings" dans le menu de gauche.
3. Cliquez sur le bouton "CI/CD" dans le menu de gauche.
4. Copiez la valeur de "Project ID" dans le champ "Promote CI Project ID" de votre pipeline Jenkins.

## Déploiement d'application avec ArgoCD

Le déploiement d'application avec ArgoCD utilise plusieurs fonctions de la shared library `sharedlib-cd`. Toutes ces fonctions prennent les paramètres suivants :

- `ARGOCD_SERVER` : l'URL du serveur ArgoCD
- `HELM_REPO_URL` : l'URL du référentiel Helm à utiliser pour le déploiement
- `HELM_CHART_PATH` : le chemin d'accès au graphique Helm à utiliser pour le déploiement
- `APP_NAME` : le nom de l'application à déployer
- `ARGOCD_PROJECT_NAME` : le nom du projet ArgoCD à utiliser pour le déploiement
- `REVISION` : la révision du graphique Helm à utiliser pour le déploiement
- `DEST_SERVER` : le nom du serveur de destination pour le déploiement
- `DEST_NAMESPACE` : l'espace de noms de destination pour le déploiement
- `HELM_VALUES_FILES` : les fichiers de valeurs Helm à utiliser pour le déploiement
- `ARGOCD_APP_FILE` : le fichier de configuration de l'application ArgoCD à utiliser

N'hésitez pas à personnaliser ce contenu en fonction de votre configuration spécifique. Si vous avez d'autres questions, je suis là pour vous aider ! 😊

import requests
import datetime
from requests.auth import HTTPBasicAuth

# Configuration d'Ansible Tower
TOWER_HOST = 'https://your-tower-instance.com'
TOWER_USERNAME = 'your-username'
TOWER_PASSWORD = 'your-password'
TOWER_VERIFY_SSL = True  # Set to False if you don't want to verify SSL certificates

# Fonction pour obtenir le token d'authentification
def get_auth_token():
    auth_url = f'{TOWER_HOST}/api/v2/authtoken/'
    auth_data = {'username': TOWER_USERNAME, 'password': TOWER_PASSWORD}
    response = requests.post(auth_url, json=auth_data, verify=TOWER_VERIFY_SSL)
    response.raise_for_status()
    return response.json().get('token', None)

# Fonction pour obtenir la date actuelle au format ISO 8601
def get_current_date():
    return datetime.datetime.now().isoformat()

# Fonction pour supprimer une organisation par ID
def delete_organization(organization_id):
    organization_url = f'{TOWER_HOST}/api/v2/organizations/{organization_id}/'
    response = requests.delete(organization_url, auth=HTTPBasicAuth(TOWER_USERNAME, TOWER_PASSWORD), verify=TOWER_VERIFY_SSL)
    response.raise_for_status()
    return response.status_code

# Obtenez le token d'authentification
auth_token = get_auth_token()

# Vérifiez si le token est disponible
if auth_token:
    # Configurez les en-têtes avec le token d'authentification
    headers = {'Authorization': f'Token {auth_token}'}

    # Récupérez la liste des organisations
    organizations_url = f'{TOWER_HOST}/api/v2/organizations/'
    organizations_response = requests.get(organizations_url, headers=headers, verify=TOWER_VERIFY_SSL)
    organizations_response.raise_for_status()
    organizations = organizations_response.json().get('results', [])

    # Récupérez la date actuelle
    current_date = get_current_date()

    # Parcourez toutes les organisations
    for organization in organizations:
        organization_id = organization['id']
        organization_name = organization['name']

        # Récupérez la liste des jobs de l'organisation
        jobs_url = f'{TOWER_HOST}/api/v2/organizations/{organization_id}/jobs/'
        jobs_response = requests.get(jobs_url, headers=headers, verify=TOWER_VERIFY_SSL)
        jobs_response.raise_for_status()
        jobs = jobs_response.json().get('results', [])

        # Vérifiez si des jobs ont été exécutés au cours des 3 derniers mois
        recent_jobs = [job for job in jobs if current_date - datetime.datetime.strptime(job['finished'], '%Y-%m-%dT%H:%M:%S.%fZ') < datetime.timedelta(days=90)]

        # Si aucun job récent, supprimez l'organisation
        if not recent_jobs:
            print(f"Aucune activité récente pour l'organisation {organization_name}. Suppression en cours...")
            delete_organization(organization_id)
            print(f"L'organisation {organization_name} a été supprimée.")
        else:
            print(f"Des activités récentes ont été détectées pour l'organisation {organization_name}. Aucune action nécessaire.")
else:
    print("Impossible d'obtenir le token d'authentification.")

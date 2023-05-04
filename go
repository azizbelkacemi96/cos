import requests
import base64

# Remplacez ces valeurs par celles de votre application
client_id = "your_client_id"
client_secret = "your_client_secret"
auth_url = "https://example.com/oauth2/token"

# Créez les données nécessaires pour l'appel API
payload = {
    "grant_type": "client_credentials",
    "scope": "your_scope"  # Remplacez "your_scope" par les scopes requis pour votre application, séparés par des espaces
}

# Encodez le client_id et le client_secret en utilisant le format Basic Auth
basic_auth = base64.b64encode(f"{client_id}:{client_secret}".encode()).decode()

# Ajoutez l'entête "Authorization" à la requête
headers = {
    "Authorization": f"Basic {basic_auth}",
    "Content-Type": "application/x-www-form-urlencoded"
}

# Effectuez l'appel API et vérifiez le statut de la réponse
response = requests.post(auth_url, data=payload, headers=headers)
response.raise_for_status()

# Récupérez le token Bearer
json_response = response.json()
access_token = json_response["access_token"]

# Utilisez le token Bearer dans vos appels API suivants
print(f"Token Bearer: {access_token}")

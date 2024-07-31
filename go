package command

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"services" // Assurez-vous que le package services est bien importé
)

// Fonction générique pour créer un pointeur vers n'importe quelle valeur
func ptr[T any](v T) *T {
	return &v
}

// Définition des options du serveur
type ServerOptions struct {
	cyberArkInstances *string
}

func TestCreateCyberArkServiceList(t *testing.T) {
	// Simuler les options de serveur
	options := &ServerOptions{
		cyberArkInstances: ptr(`[
			{
				"MUT": {
					"appID": "server1",
					"apiURL": "https://api.server1.com",
					"apiClientCert": "-----BEGIN CERTIFICATE-----\nMIIC5jCCAc6gAwIBAgIJAK4f7hx/40YwMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0ZXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTUwNjAxMTMxNDA5WhcNMTYwNjAxMTMxNDA5\nWjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwY\nSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\nMIIBCgKCAQEAq/T41G/x3h8C4e19PQJf+f7WqFVYW4b/xz36f66t0+/KyMJqf7Ud\nXDJRs93jFjz/S492Fu1P2EZ9D1Ou/KL15Qq6V/Ze95Zn9y8b/7J978X13QH132Nt\n-----END CERTIFICATE-----",
					"apiClientKey": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC/G41v0zKuU7vW\n-----END PRIVATE KEY-----",
					"sdkPath": ""
				},
				"CIS": {
					"appID": "server2",
					"apiURL": "https://api.server2.com",
					"apiClientCert": "-----BEGIN CERTIFICATE-----\nMIIC5jCCAc6gAwIBAgIJAK4f7hx/40YwMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0ZXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTUwNjAxMTMxNDA5WhcNMTYwNjAxMTMxNDA5\nWjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwY\nSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\nMIIBCgKCAQEAq/T41G/x3h8C4e19PQJf+f7WqFVYW4b/xz36f66t0+/KyMJqf7Ud\nXDJRs93jFjz/S492Fu1P2EZ9D1Ou/KL15Qq6V/Ze95Zn9y8b/7J978X13QH132Nt\n-----END CERTIFICATE-----",
					"apiClientKey": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC/G41v0zKuU7vW\n-----END PRIVATE KEY-----",
					"sdkPath": ""
				}
			}
		]`),
	}

	// Simuler la création des services
	var cyberArkMapping []map[string]map[string]string
	err := json.Unmarshal([]byte(*options.cyberArkInstances), &cyberArkMapping)
	assert.NoError(t, err)

	// Création de la liste de services CyberArk
	cyberArkServiceList := make(map[string]services.CyberArkServiceInterface)

	// Parcours de chaque instance CyberArk
	for _, cyberArkInstance := range cyberArkMapping {
		for key, instance := range cyberArkInstance {
			// Utilisation des clés prédéfinies "MUT" et "CIS"
			if instance["sdkPath"] == "" {
				// Création du service API
				service, err := services.NewCyberArkAPIService(
					instance["appID"],
					instance["apiURL"],
					instance["apiClientCert"],
					instance["apiClientKey"],
				)
				if err != nil {
					t.Fatalf("Failed to create CyberArk APIService for %s: %v", key, err)
				}
				// Ajout du service à la liste correspondante
				cyberArkServiceList[key] = service
			} else {
				// Création du service CLI
				service, err := services.NewCyberArkCLIService(
					instance["appID"],
					instance["sdkPath"],
				)
				if err != nil {
					t.Fatalf("Failed to create CyberArk CLIService for %s: %v", key, err)
				}
				// Ajout du service à la liste correspondante
				cyberArkServiceList[key] = service
			}
		}
	}

	// Vérification que les services ont été créés correctement
	assert.NotNil(t, cyberArkServiceList["MUT"])
	assert.NotNil(t, cyberArkServiceList["CIS"])

	// Log des services créés pour inspection
	for key, service := range cyberArkServiceList {
		t.Logf("Key: %s, Service: %v", key, service)
	}

	// Exemple d'appel de méthode pour le service MUT
	c, err := someContextCreationFunction()
	if err != nil {
		t.Fatalf("Failed to create context: %v", err)
	}
	serverPassword, err := cyberArkServiceList["MUT"].GetVaultPassword(c)
	if err != nil {
		t.Fatalf("Failed to get vault password: %v", err)
	}

	// Utiliser le serverPassword
	fmt.Println("Server Password:", serverPassword)
}

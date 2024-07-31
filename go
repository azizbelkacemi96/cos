package command

import (
	"encoding/json"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"

	"services" // Assurez-vous que le package services est bien importé
)

func ptr[T any](v T) *T {
	return &v
}

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
	for idx, cyberArkInstance := range cyberArkMapping {
		// Construction d'une clé unique pour chaque service
		for key, instance := range cyberArkInstance {
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
				uniqueKey := fmt.Sprintf("%s-%d", key, idx)
				cyberArkServiceList[uniqueKey] = service
			} else {
				// Création du service CLI
				service, err := services.NewCyberArkCLIService(
					instance["appID"],
					instance["sdkPath"],
				)
				if err != nil {
					t.Fatalf("Failed to create CyberArk CLIService for %s: %v", key, err)
				}
				uniqueKey := fmt.Sprintf("%s-%d", key, idx)
				cyberArkServiceList[uniqueKey] = service
			}
		}
	}

	// Vérification que les services ont été créés correctement
	assert.NotNil(t, cyberArkServiceList["MUT-0"])
	assert.NotNil(t, cyberArkServiceList["CIS-0"])

	// Log des services créés pour inspection
	t.Logf("cyberArkServiceList: %+v", cyberArkServiceList)
}

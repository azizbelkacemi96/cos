package command

import (
	"encoding/json"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func ptr[T any](v T) *T {
	return &v
}

func TestCreateCyberArkServiceList(t *testing.T) {
	// Simuler les options de serveur
	options := &ServerOptions{
		cyberArkInstances: ptr(`[
			{
				"MUT": {
					"appID": "server1",
					"apiURL": "https://api.server1.com",
					"apiClientCert": "-----BEGIN CERTIFICATE-----\nMIIC5jCCAc6gAwIBAgIJAK4f7hx/40YwMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTUwNjAxMTMxNDA5WhcNMTYwNjAxMTMxNDA5\nWjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwY\nSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\nMIIBCgKCAQEAq/T41G/x3h8C4e19PQJf+f7WqFVYW4b/xz36f66t0+/KyMJqf7Ud\nXDJRs93jFjz/S492Fu1P2EZ9D1Ou/KL15Qq6V/Ze95Zn9y8b/7J978X13QH132Nt\n-----END CERTIFICATE-----",
					"apiClientKey": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC/G41v0zKuU7vW\n-----END PRIVATE KEY-----",
					"sdkPath": ""
				},
				"CIS": {
					"appID": "server2",
					"apiURL": "https://api.server2.com",
					"apiClientCert": "-----BEGIN CERTIFICATE-----\nMIIC5jCCAc6gAwIBAgIJAK4f7hx/40YwMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTUwNjAxMTMxNDA5WhcNMTYwNjAxMTMxNDA5\nWjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwY\nSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\nMIIBCgKCAQEAq/T41G/x3h8C4e19PQJf+f7WqFVYW4b/xz36f66t0+/KyMJqf7Ud\nXDJRs93jFjz/S492Fu1P2EZ9D1Ou/KL15Qq6V/Ze95Zn9y8b/7J978X13QH132Nt\n-----END CERTIFICATE-----",
					"apiClientKey": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC/G41v0zKuU7vW\n-----END PRIVATE KEY-----",
					"sdkPath": ""
				}
			},
			{
				"MUT": {
					"appID": "server3",
					"apiURL": "",
					"apiClientCert": "",
					"apiClientKey": "",
					"sdkPath": "/path/to/sdk"
				},
				"CIS": {
					"appID": "server4",
					"apiURL": "",
					"apiClientCert": "",
					"apiClientKey": "",
					"sdkPath": "/path/to/sdk"
				}
			}
		]`),
	}

	// Simuler le logger
	zapLogger, _ := zap.NewDevelopment()
	logger := &ZapLogger{Logger: zapLogger}
	e := echo.New()
	e.Logger = logger.Output()

	// Simuler la création des services
	var cyberArkMapping []map[string]map[string]string
	err := json.Unmarshal([]byte(*options.cyberArkInstances), &cyberArkMapping)
	assert.NoError(t, err)

	cyberArkServiceList := make(map[string]services.CyberArkServiceInterface)
	for _, cyberArkInstance := range cyberArkMapping {
		if cyberArkInstance["MUT"]["sdkPath"] == "" {
			service, err := services.NewCyberArkAPIService(cyberArkInstance["MUT"]["appID"], cyberArkInstance["MUT"]["apiURL"], cyberArkInstance["MUT"]["apiClientCert"], cyberArkInstance["MUT"]["apiClientKey"])
			if err != nil {
				t.Fatalf("Failed to create CyberArk APIService for MUT: %v", err)
			}
			cyberArkServiceList["MUT"] = service
		} else {
			service, err := services.NewCyberArkCLIService(cyberArkInstance["MUT"]["appID"], cyberArkInstance["MUT"]["sdkPath"])
			if err != nil {
				t.Fatalf("Failed to create CyberArk CLIService for MUT: %v", err)
			}
			cyberArkServiceList["MUT"] = service
		}

		if cyberArkInstance["CIS"]["sdkPath"] == "" {
			service, err := services.NewCyberArkAPIService(cyberArkInstance["CIS"]["appID"], cyberArkInstance["CIS"]["apiURL"], cyberArkInstance["CIS"]["apiClientCert"], cyberArkInstance["CIS"]["apiClientKey"])
			if err != nil {
				t.Fatalf("Failed to create CyberArk APIService for CIS: %v", err)
			}
			cyberArkServiceList["CIS"] = service
		} else {
			service, err := services.NewCyberArkCLIService(cyberArkInstance["CIS"]["appID"], cyberArkInstance["CIS"]["sdkPath"])
			if err != nil {
				t.Fatalf("Failed to create CyberArk CLIService for CIS: %v", err)
			}
			cyberArkServiceList["CIS"] = service
		}
	}

	// Vérifier que les services ont été créés correctement
	assert.NotNil(t, cyberArkServiceList["MUT"])
	assert.NotNil(t, cyberArkServiceList["CIS"])
}

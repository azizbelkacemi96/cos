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
					"apiClientCert": "/path/to/cert.pem",
					"apiClientKey": "/path/to/key.pem",
					"sdkPath": ""
				},
				"CIS": {
					"appID": "server2",
					"apiURL": "https://api.server2.com",
					"apiClientCert": "/path/to/cert.pem",
					"apiClientKey": "/path/to/key.pem",
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
	logger, _ := zap.NewDevelopment()
	e := echo.New()
	e.Logger = logger

	// Simuler la création des services
	var cyberArkMapping []map[string]map[string]string
	err := json.Unmarshal([]byte(*options.cyberArkInstances), &cyberArkMapping)
	assert.NoError(t, err)

	cyberArkServiceList := make(map[string]services.CyberArkServiceInterface)
	for _, cyberArkInstance := range cyberArkMapping {
		if cyberArkInstance["MUT"]["sdkPath"] == "" || cyberArkInstance["CIS"]["sdkPath"] == "" {
			cyberArkServiceList["MUT"], err = services.NewCyberArkAPIService(cyberArkInstance["MUT"]["appID"], cyberArkInstance["MUT"]["apiURL"], cyberArkInstance["MUT"]["apiClientCert"], cyberArkInstance["MUT"]["apiClientKey"])
			cyberArkServiceList["CIS"], err = services.NewCyberArkAPIService(cyberArkInstance["CIS"]["appID"], cyberArkInstance["CIS"]["apiURL"], cyberArkInstance["CIS"]["apiClientCert"], cyberArkInstance["CIS"]["apiClientKey"])
		} else {
			cyberArkServiceList["MUT"], err = services.NewCyberArkCLIService(cyberArkInstance["MUT"]["appID"], cyberArkInstance["MUT"]["sdkPath"])
			cyberArkServiceList["CIS"], err = services.NewCyberArkCLIService(cyberArkInstance["CIS"]["appID"], cyberArkInstance["CIS"]["sdkPath"])
		}
		assert.NoError(t, err)
	}

	// Vérifier que les services ont été créés correctement
	assert.NotNil(t, cyberArkServiceList["MUT"])
	assert.NotNil(t, cyberArkServiceList["CIS"])
}

cyberArkServiceList := make(map[string]services.CyberArkServiceInterface)
for _, cyberArkInstance := range cyberArkMapping {
	var err error
	if cyberArkInstance["MUT"]["sdkPath"] == "" || cyberArkInstance["CIS"]["sdkPath"] == "" {
		cyberArkServiceList["MUT"], err = services.NewCyberArkAPIService(cyberArkInstance["MUT"]["appID"], cyberArkInstance["MUT"]["apiURL"], cyberArkInstance["MUT"]["apiClientCert"], cyberArkInstance["MUT"]["apiClientKey"])
		if err != nil {
			assert.NoError(t, err)
			continue
		}
		cyberArkServiceList["CIS"], err = services.NewCyberArkAPIService(cyberArkInstance["CIS"]["appID"], cyberArkInstance["CIS"]["apiURL"], cyberArkInstance["CIS"]["apiClientCert"], cyberArkInstance["CIS"]["apiClientKey"])
	} else {
		cyberArkServiceList["MUT"], err = services.NewCyberArkCLIService(cyberArkInstance["MUT"]["appID"], cyberArkInstance["MUT"]["sdkPath"])
		if err != nil {
			assert.NoError(t, err)
			continue
		}
		cyberArkServiceList["CIS"], err = services.NewCyberArkCLIService(cyberArkInstance["CIS"]["appID"], cyberArkInstance["CIS"]["sdkPath"])
	}
	assert.NoError(t, err)
}

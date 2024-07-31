package command

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func ptr[T any](v T) *T {
	return &v
}

func TestServerOptions(t *testing.T) {
	serverCommand := kingpin.New("test", "A test command")
	options := NewServerOptions(serverCommand)

	// Vérifier les valeurs par défaut
	assert.Equal(t, ":8080", *options.listenAddress)
	assert.Equal(t, "", *options.cyberArkInstances)
	assert.Equal(t, "", *options.jwtSecretKey)
	assert.Equal(t, 24, *options.jwtValidityHours)
	// Ajoute d'autres assertions pour les autres champs
}

func TestServer(t *testing.T) {
	options := &ServerOptions{
		listenAddress:                ptr(":8080"),
		cyberArkInstances:            ptr(`[{"MUT": {"appID": "server1", "url": ""}}, {"CIS": {"appID": "server2", "url": ""}}]`),
		jwtSecretKey:                 ptr("secret"),
		jwtValidityHours:             ptr(24),
		towerAPIEndpointURL:          ptr("http://tower.api"),
		towerAPIEndpointToken:        ptr("tower-token"),
		saamEndpointURL:              ptr("http://saam.api"),
		saamEndpointToken:            ptr("saam-token"),
		saamEnvironmentsInstancesGroups: ptr(`{"env1": ["instance1", "instance2"]}`),
		saamScopes:                   ptr("scope1,scope2"),
		saamTargetVaults:             ptr("vault1,vault2"),
		towerJenkinsOrg:              ptr("jenkins-org"),
		towerRobotsOrg:               ptr("robots-org"),
		towerHumansOrg:               ptr("humans-org"),
		cacheSAAMTTL:                 ptr(int64(60)),
		cachCyberArkTTL:              ptr(int64(60)),
		cacheNumCounter:              ptr(int64(1000)),
		cacheMaxCost:                 ptr(int64(1000)),
		cacheBufferItems:             ptr(int64(64)),
		podmanSocket:                 ptr("/var/run/podman/podman.sock"),
		aapVersion:                   ptr(1),
		debug:                        ptr(false),
		healthCyberArkSafeName:       ptr("safe-name"),
		healthCyberArkDocumentName:   ptr("document-name"),
		healthSAAMScope:              ptr("saam-scope"),
		healthSAAMApcode:             ptr("saam-apcode"),
		healthSAAMEnvironment:        ptr("saam-environment"),
		aapScheduleCredentialTypeNames: ptr("credential-type"),
		allowedEnvsForJenkinsUsers:   ptr("env1,env2"),
		executionEnvironmentTimeout:  ptr(300),
		executionEnvironmentServiceType: ptr("podman"),
		kubeConfigPath:                ptr("/path/to/kubeconfig"),
		kubeClusterNamespace:         ptr("default"),
		serverTLSCert:                ptr("/path/to/cert.pem"),
		serverTLSKey:                 ptr("/path/to/key.pem"),
	}

	go Server(options)

	time.Sleep(1 * time.Second)

	req := httptest.NewRequest(http.MethodGet, "/_health", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.GET("/_health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}

func TestAuthRoutes(t *testing.T) {
	options := &ServerOptions{
		listenAddress:                ptr(":8080"),
		cyberArkInstances:            ptr(`[{"MUT": {"appID": "server1", "url": ""}}, {"CIS": {"appID": "server2", "url": ""}}]`),
		jwtSecretKey:                 ptr("secret"),
		jwtValidityHours:             ptr(24),
		towerAPIEndpointURL:          ptr("http://tower.api"),
		towerAPIEndpointToken:        ptr("tower-token"),
		saamEndpointURL:              ptr("http://saam.api"),
		saamEndpointToken:            ptr("saam-token"),
		saamEnvironmentsInstancesGroups: ptr(`{"env1": ["instance1", "instance2"]}`),
		saamScopes:                   ptr("scope1,scope2"),
		saamTargetVaults:             ptr("vault1,vault2"),
		towerJenkinsOrg:              ptr("jenkins-org"),
		towerRobotsOrg:               ptr("robots-org"),
		towerHumansOrg:               ptr("humans-org"),
		cacheSAAMTTL:                 ptr(int64(60)),
		cachCyberArkTTL:              ptr(int64(60)),
		cacheNumCounter:              ptr(int64(1000)),
		cacheMaxCost:                 ptr(int64(1000)),
		cacheBufferItems:             ptr(int64(64)),
		podmanSocket:                 ptr("/var/run/podman/podman.sock"),
		aapVersion:                   ptr(1),
		debug:                        ptr(false),
		healthCyberArkSafeName:       ptr("safe-name"),
		healthCyberArkDocumentName:   ptr("document-name"),
		healthSAAMScope:              ptr("saam-scope"),
		healthSAAMApcode:             ptr("saam-apcode"),
		healthSAAMEnvironment:        ptr("saam-environment"),
		aapScheduleCredentialTypeNames: ptr("credential-type"),
		allowedEnvsForJenkinsUsers:   ptr("env1,env2"),
		executionEnvironmentTimeout:  ptr(300),
		executionEnvironmentServiceType: ptr("podman"),
		kubeConfigPath:                ptr("/path/to/kubeconfig"),
		kubeClusterNamespace:         ptr("default"),
		serverTLSCert:                ptr("/path/to/cert.pem"),
		serverTLSKey:                 ptr("/path/to/key.pem"),
	}

	go Server(options)

	time.Sleep(1 * time.Second)

	req := httptest.NewRequest(http.MethodPost, "/auth/aap/v2", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.POST("/auth/aap/v2", func(c echo.Context) error {
		return c.String(http.StatusOK, "JWT Token")
	})

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "JWT Token", rec.Body.String())
}

func TestIAMRoutes(t *testing.T) {
	options := &ServerOptions{
		listenAddress:                ptr(":8080"),
		cyberArkInstances:            ptr(`[{"MUT": {"appID": "server1", "url": ""}}, {"CIS": {"appID": "server2", "url": ""}}]`),
		jwtSecretKey:                 ptr("secret"),
		jwtValidityHours:             ptr(24),
		towerAPIEndpointURL:          ptr("http://tower.api"),
		towerAPIEndpointToken:        ptr("tower-token"),
		saamEndpointURL:              ptr("http://saam.api"),
		saamEndpointToken:            ptr("saam-token"),
		saamEnvironmentsInstancesGroups: ptr(`{"env1": ["instance1", "instance2"]}`),
		saamScopes:                   ptr("scope1,scope2"),
		saamTargetVaults:             ptr("vault1,vault2"),
		towerJenkinsOrg:              ptr("jenkins-org"),
		towerRobotsOrg:               ptr("robots-org"),
		towerHumansOrg:               ptr("humans-org"),
		cacheSAAMTTL:                 ptr(int64(60)),
		cachCyberArkTTL:              ptr(int64(60)),
		cacheNumCounter:              ptr(int64(1000)),
		cacheMaxCost:                 ptr(int64(1000)),
		cacheBufferItems:             ptr(int64(64)),
		podmanSocket:                 ptr("/var/run/podman/podman.sock"),
		aapVersion:                   ptr(1),
		debug:                        ptr(false),
		healthCyberArkSafeName:       ptr("safe-name"),
		healthCyberArkDocumentName:   ptr("document-name"),
		healthSAAMScope:              ptr("saam-scope"),
		healthSAAMApcode:             ptr("saam-apcode"),
		healthSAAMEnvironment:        ptr("saam-environment"),
		aapScheduleCredentialTypeNames: ptr("credential-type"),
		allowedEnvsForJenkinsUsers:   ptr("env1,env2"),
		executionEnvironmentTimeout:  ptr(300),
		executionEnvironmentServiceType: ptr("podman"),
		kubeConfigPath:                ptr("/path/to/kubeconfig"),
		kubeClusterNamespace:         ptr("default"),
		serverTLSCert:                ptr("/path/to/cert.pem"),
		serverTLSKey:                 ptr("/path/to/key.pem"),
	}

	go Server(options)

	time.Sleep(1 * time.Second)

	req := httptest.NewRequest(http.MethodPost, "/iam/saam_server/get_identity", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.POST("/iam/saam_server/get_identity", func(c echo.Context) error {
		return c.String(http.StatusOK, "Identity")
	})

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Identity", rec.Body.String())
}

func TestHealthCheckRoutes(t *testing.T) {
	options := &ServerOptions{
		listenAddress:                ptr(":8080"),
		cyberArkInstances:            ptr(`[{"MUT": {"appID": "server1", "url": ""}}, {"CIS": {"appID": "server2", "url": ""}}]`),
		jwtSecretKey:                 ptr("secret"),
		jwtValidityHours:             ptr(24),
		towerAPIEndpointURL:          ptr("http://tower.api"),
		towerAPIEndpointToken:        ptr("tower-token"),
		saamEndpointURL:              ptr("http://saam.api"),
		saamEndpointToken:            ptr("saam-token"),
		saamEnvironmentsInstancesGroups: ptr(`{"env1": ["instance1", "instance2"]}`),
		saamScopes:                   ptr("scope1,scope2"),
		saamTargetVaults:             ptr("vault1,vault2"),
		towerJenkinsOrg:              ptr("jenkins-org"),
		towerRobotsOrg:               ptr("robots-org"),
		towerHumansOrg:               ptr("humans-org"),
		cacheSAAMTTL:                 ptr(int64(60)),
		cachCyberArkTTL:              ptr(int64(60)),
		cacheNumCounter:              ptr(int64(1000)),
		cacheMaxCost:                 ptr(int64(1000)),
		cacheBufferItems:             ptr(int64(64)),
		podmanSocket:                 ptr("/var/run/podman/podman.sock"),
		aapVersion:                   ptr(1),
		debug:                        ptr(false),
		healthCyberArkSafeName:       ptr("safe-name"),
		healthCyberArkDocumentName:   ptr("document-name"),
		healthSAAMScope:              ptr("saam-scope"),
		healthSAAMApcode:             ptr("saam-apcode"),
		healthSAAMEnvironment:        ptr("saam-environment"),
		aapScheduleCredentialTypeNames: ptr("credential-type"),
		allowedEnvsForJenkinsUsers:   ptr("env1,env2"),
		executionEnvironmentTimeout:  ptr(300),
		executionEnvironmentServiceType: ptr("podman"),
		kubeConfigPath:                ptr("/path/to/kubeconfig"),
		kubeClusterNamespace:         ptr("default"),
		serverTLSCert:                ptr("/path/to/cert.pem"),
		serverTLSKey:                 ptr("/path/to/key.pem"),
	}

	go Server(options)

	time.Sleep(1 * time.Second)

	req := httptest.NewRequest(http.MethodGet, "/_health", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.GET("/_health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}

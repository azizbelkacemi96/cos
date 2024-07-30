package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/stretchr/testify/mock"
    "go.uber.org/zap"
)

type MockExecutionEnvironmentService struct {
    mock.Mock
}

type MockTowerService struct {
    mock.Mock
}

type MockSaamService struct {
    mock.Mock
}

type MockCyberArkService struct {
    mock.Mock
}

func TestHandlers(t *testing.T) {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    e := echo.New()

    options := &Options{
        jwtValidityHours: new(int),
        jwtSecretKey:     new(string),
        towerJenkinsOrg:  new(string),
        towerRobotsOrg:   new(string),
        towerHumansOrg:   new(string),
        aapVersion:       new(string),
        aapScheduleCredentialTypeNames: new(string),
        cacheSAAMTTL:     new(int),
        cachCyberArkTTL:  new(int),
    }

    *options.jwtValidityHours = 24
    *options.jwtSecretKey = "secret"
    *options.towerJenkinsOrg = "jenkinsOrg"
    *options.towerRobotsOrg = "robotsOrg"
    *options.towerHumansOrg = "humansOrg"
    *options.aapVersion = "v2"
    *options.aapScheduleCredentialTypeNames = "scheduleCredentialTypeNames"
    *options.cacheSAAMTTL = 60
    *options.cachCyberArkTTL = 60

    executionEnvironmentService := new(MockExecutionEnvironmentService)
    towerService := new(MockTowerService)
    saamService := new(MockSaamService)
    cyberArkServiceList := []MockCyberArkService{*new(MockCyberArkService)}

    jwtAuth := auth.NewJWTAuth(
        time.Duration(*options.jwtValidityHours)*time.Hour,
        *options.jwtSecretKey,
        *options.towerJenkinsOrg,
        *options.towerRobotsOrg,
        *options.towerHumansOrg,
    )

    jwtConfig := middleware.JWTConfig{
        Claims:     &auth.JwtTowerJobIdentityClaims{},
        SigningKey: []byte(*options.jwtSecretKey),
    }

    authHandler := handlers.NewAuthHandler(executionEnvironmentService, towerService, saamService, jwtAuth, *options.aapVersion, *options.aapScheduleCredentialTypeNames)

    authGroup := e.Group("/auth")
    authGroup.POST("/aap/v2", authHandler.GenerateJWTAAP)

    iamHandler := handlers.NewIAMHandler(cyberArkServiceList, saamService, *options.cacheSAAMTTL, *options.cachCyberArkTTL, cacheManager, saamEnvMapping, allowedEnvsForJenkinsUsersParsed)

    iamGroup := e.Group("/iam")
    iamGroup.Use(middleware.JWTWithConfig(jwtConfig))
    iamGroup.POST("/saam_server/get_identity", iamHandler.GetServerIdentity)

    // Test the /auth/aap/v2 endpoint
    req := httptest.NewRequest(http.MethodPost, "/auth/aap/v2", strings.NewReader(`{"username":"test","password":"test"}`))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, authHandler.GenerateJWTAAP(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Contains(t, rec.Body.String(), "token")
    }

    // Test the /iam/saam_server/get_identity endpoint
    req = httptest.NewRequest(http.MethodPost, "/iam/saam_server/get_identity", strings.NewReader(`{"server_id":"test_server"}`))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    req.Header.Set(echo.HeaderAuthorization, "Bearer "+generateTestJWT(*options.jwtSecretKey))
    rec = httptest.NewRecorder()
    c = e.NewContext(req, rec)

    if assert.NoError(t, iamHandler.GetServerIdentity(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Contains(t, rec.Body.String(), "identity")
    }
}

func generateTestJWT(secret string) string {
    // Generate a test JWT token
    return "test_jwt_token"
}

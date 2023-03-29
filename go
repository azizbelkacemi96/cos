package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockSaamService struct {
	success bool
}

type mockTowerService struct {
	success bool
}

func (m *mockSaamService) GetHealthCheck(c echo.Context, scope, apcode, environement string) (bool, error) {
	if m.success {
		return true, nil
	}
	return false, errors.New("saam error")
}

func (m *mockTowerService) GetHealthCheck(c echo.Context) (bool, error) {
	if m.success {
		return true, nil
	}
	return false, errors.New("tower error")
}

func TestGetHealthCheck_Success(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &HealthCheckHandler{
		SaamService:  &mockSaamService{success: true},
		TowerService: &mockTowerService{success: true},
		// ... autres champs
	}

	err := h.GetHealthCheck(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetHealthCheck_Failure(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &HealthCheckHandler{
		SaamService:  &mockSaamService{success: false},
		TowerService: &mockTowerService{success: true},
		// ... autres champs
	}

	err := h.GetHealthCheck(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
}

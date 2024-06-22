package edicionEquipoController

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEdicionEquipoInsert(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/edicion-equipo", EdicionEquipoInsert)

	// Mock request body
	requestBody := []byte(`{"example_field": "example_value"}`)

	// Create a mock HTTP request
	req, err := http.NewRequest("POST", "/edicion-equipo", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Optionally, check the response body or headers here
	// Example: Assert that the response body contains expected data
	assert.Contains(t, w.Body.String(), "example_value")
}

func TestGetEdicionEquipos(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/edicion-equipos/:id_edicion_torneo", GetEdicionEquipos)

	// Set up a mock HTTP request
	idEdicionTorneo := "123" // Example ID
	req, err := http.NewRequest("GET", "/edicion-equipos/"+idEdicionTorneo, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Optionally, check the response body or headers here
	// Example: Assert that the response body is not empty
	assert.NotEmpty(t, w.Body.String())
}

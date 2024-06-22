package resultadoController

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Función para inicializar el router de Gin para los tests
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/resultado", ResultadoInsert)
	router.GET("/resultados", GetResultados)
	router.GET("/resultado/:id_edicion_torneos", GetResultadoByIdEdicionTorneo)
	return router
}

// Test para ResultadoInsert
func TestResultadoInsert(t *testing.T) {
	router := setupRouter()

	// Mock request body
	requestBody := []byte(`{"example_field": "example_value"}`)

	// Create a mock HTTP request
	req, err := http.NewRequest("POST", "/resultado", bytes.NewBuffer(requestBody))
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

// Test para GetResultados
func TestGetResultados(t *testing.T) {
	router := setupRouter()

	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/resultados", nil)
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

// Test para GetResultadoByIdEdicionTorneo
func TestGetResultadoByIdEdicionTorneo(t *testing.T) {
	router := setupRouter()

	// Set up a mock HTTP request
	idEdicionTorneo := "123" // Example ID
	req, err := http.NewRequest("GET", "/resultado/"+idEdicionTorneo, nil)
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

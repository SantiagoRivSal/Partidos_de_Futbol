package equipoController

import (
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
	router.GET("/equipos", GetEquipos)
	router.GET("/equipos/:id_pais", GetEquiposByIdPais)
	router.GET("/equipo/:id", GetEquipoById)
	return router
}

// Test para GetEquipos
func TestGetEquipos(t *testing.T) {
	router := setupRouter()

	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/equipos", nil)
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

// Test para GetEquiposByIdPais
func TestGetEquiposByIdPais(t *testing.T) {
	router := setupRouter()

	// Set up a mock HTTP request
	idPais := "123" // Example ID
	req, err := http.NewRequest("GET", "/equipos/"+idPais, nil)
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

// Test para GetEquipoById
func TestGetEquipoById(t *testing.T) {
	router := setupRouter()

	// Set up a mock HTTP request
	idEquipo := "456" // Example ID
	req, err := http.NewRequest("GET", "/equipo/"+idEquipo, nil)
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

package torneoController

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/dto"
	"backend/services"
	"backend/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTorneos(t *testing.T) {
	// Mock del servicio para simular la obtención de Torneos
	services.TorneoService = &MockTorneoService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.GET("/torneos", GetTorneos)

	// Realizar una solicitud HTTP GET simulada a /Torneos
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/torneos", nil)
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto TorneosDto
	var responseDto dto.TorneosDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos de confederaciones simulados
	assert.NotEmpty(t, responseDto)
}

// MockPaisService es una implementación de PaisService para usar en las pruebas
type MockTorneoService struct{}

// GetTorneos simula la obtención de Torneos
func (m *MockTorneoService) GetTorneos() (dto.TorneosDto, errors.ApiError) {
	// Simular una lista de Torneos
	torneos := []dto.TorneoDto{
		{Id: 1, Nombre: "Torneo 1", Logo: "logo1.png", IdConfederacion: 1},
		{Id: 2, Nombre: "Torneo 2", Logo: "logo2.png", IdConfederacion: 1},
	}

	return torneos, nil
}

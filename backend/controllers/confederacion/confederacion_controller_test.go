package confederacionController

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

func TestGetConfederaciones(t *testing.T) {
	// Mock del servicio para simular la obtención de confederaciones
	services.ConfederacionService = &MockConfederacionService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.GET("/confederaciones", GetConfederaciones)

	// Realizar una solicitud HTTP GET simulada a /confederaciones
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/confederaciones", nil)
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto ConfederacionesDto
	var responseDto dto.ConfederacionesDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos de confederaciones simulados
	assert.NotEmpty(t, responseDto)
}

// MockConfederacionService es una implementación de ConfederacionService para usar en las pruebas
type MockConfederacionService struct{}

// GetConfederaciones simula la obtención de Confederaciones
func (m *MockConfederacionService) GetConfederaciones() (dto.ConfederacionesDto, errors.ApiError) {
	// Simular una lista de confederaciones
	confederaciones := []dto.ConfederacionDto{
		{Id: 1, Nombre: "Confederación 1", Logo: "logo1.png"},
		{Id: 2, Nombre: "Confederación 2", Logo: "logo2.png"},
	}

	return confederaciones, nil
}

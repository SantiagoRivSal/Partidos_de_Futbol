package confederacionController

import (
	"backend/dto"
	"backend/services"
	e "backend/utils/errors"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetConfederaciones(t *testing.T) {
	// Mock del servicio para simular la obtención de Confederaciones
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

	// Verificar que la respuesta contiene datos de productos simulados
	assert.NotEmpty(t, responseDto)

}

// MockConfederacionService es una implementación de ConfederacionService para usar en las pruebas
type MockConfederacionService struct{}

// GetProducts simula la obtención de productos
func (m *MockConfederacionService) GetConfederaciones() (dto.ConfederacionesDto, e.ApiError) {
	// Simular una lista de productos
	confederaciones := []dto.ConfederacionDto{
		{Id: 1, Nombre: "Confederación 1", Logo: "logo1.png"},
		{Id: 2, Nombre: "Confederación 2", Logo: "logo2.png"},
	}

	return confederaciones, nil
}

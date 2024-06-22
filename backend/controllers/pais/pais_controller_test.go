package paisController

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

func TestGetPaises(t *testing.T) {
	// Mock del servicio para simular la obtención de Paises
	services.PaisService = &MockPaisService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.GET("/paises", GetPaises)

	// Realizar una solicitud HTTP GET simulada a /paises
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/paises", nil)
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto PaisesDto
	var responseDto dto.PaisesDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos de productos simulados
	assert.NotEmpty(t, responseDto)

}

// MockPaisService es una implementación de PaisService para usar en las pruebas
type MockPaisService struct{}

// GetProducts simula la obtención de productos
func (m *MockPaisService) GetPaises() (dto.PaisesDto, e.ApiError) {
	// Simular una lista de productos
	Paises := []dto.PaisDto{
		{Id: 1, Nombre: "Pais 1", Bandera: "bandera1.png"},
		{Id: 2, Nombre: "Pais 2", Bandera: "bandera2.png"},
	}

	return Paises, nil
}

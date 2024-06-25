package resultadoController

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/dto"
	"backend/services"
	"backend/utils/errors"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetResultados(t *testing.T) {
	// Mock del servicio para simular la obtención de confederaciones
	services.ResultadoService = &MockResultadoService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.GET("/resultados", GetResultados)

	// Realizar una solicitud HTTP GET simulada a /confederaciones
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/resultados", nil)
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto ConfederacionesDto
	var responseDto dto.ResultadosDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos de confederaciones simulados
	assert.NotEmpty(t, responseDto)
}

func TestInsertResultados(t *testing.T) {
	// Mock del servicio para simular la inserción de resultados
	services.ResultadoService = &MockResultadoService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.POST("/resultado", ResultadoInsert)

	// Crear un objeto ProductDto simulado para la solicitud JSON
	resultadoDto := dto.ResultadoDto{Id: 1, IdEdicionTorneo: 2023, Campeon: 10, Subcampeon: 20}

	// Convertir el objeto ProductDto en JSON
	payload, err := json.Marshal(resultadoDto)
	assert.Nil(t, err)

	// Realizar una solicitud HTTP POST simulada a /product
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/resultado", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusCreated, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto ProductDto
	var responseDto dto.ResultadoDto
	err = json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos del producto simulado
	assert.Equal(t, resultadoDto.Id, responseDto.Id)
	assert.Equal(t, resultadoDto.IdEdicionTorneo, responseDto.IdEdicionTorneo)
	assert.Equal(t, resultadoDto.Campeon, responseDto.Campeon)
	assert.Equal(t, resultadoDto.Subcampeon, responseDto.Subcampeon)
}

/*func TestGetResultadoByIdEdicionTorneo (t *testing.T) {

}*/

// MockConfederacionService es una implementación de ConfederacionService para usar en las pruebas
type MockResultadoService struct{}

// GetResultadoByIdEdicionTorneo implements services.resultadoServiceInterface.
func (m *MockResultadoService) GetResultadoByIdEdicionTorneo(IdEdicionTorneo int) (dto.ResultadoDto, errors.ApiError) {
	panic("unimplemented")
}

// InsertResultados implements services.resultadoServiceInterface.
func (m *MockResultadoService) InsertResultados(resultadoDto dto.ResultadoDto) (dto.ResultadoDto, errors.ApiError) {
	return resultadoDto, nil
}

// GetConfederaciones simula la obtención de Confederaciones
func (m *MockResultadoService) GetResultados() (dto.ResultadosDto, errors.ApiError) {
	// Simular una lista de confederaciones
	resultados := []dto.ResultadoDto{
		{Id: 1, IdEdicionTorneo: 2023, Campeon: 10, Subcampeon: 20},
		{Id: 2, IdEdicionTorneo: 2024, Campeon: 30, Subcampeon: 40},
	}

	return resultados, nil
}

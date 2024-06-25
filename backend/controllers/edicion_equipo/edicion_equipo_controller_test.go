package edicionEquipoController

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

func TestEdicionEquipoInsert(t *testing.T) {
	// Mock del servicio para simular la inserción de resultados
	services.EdicionEquipoService = &MockEdicionEquipoService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.POST("/equipoxedicion", EdicionEquipoInsert)

	// Crear un objeto ProductDto simulado para la solicitud JSON
	edicionEquipoDto := dto.EdicionEquipoDto{
		Id: 1, IdEdicionTorneo: 1, IdEquipo: 1,
	}

	// Convertir el objeto ProductDto en JSON
	payload, err := json.Marshal(edicionEquipoDto)
	assert.Nil(t, err)

	// Realizar una solicitud HTTP POST simulada a /product
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/equipoxedicion", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusCreated, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto ProductDto
	var responseDto dto.EdicionEquipoDto
	err = json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos del producto simulado
	assert.Equal(t, edicionEquipoDto.Id, responseDto.Id)
	assert.Equal(t, edicionEquipoDto.IdEdicionTorneo, responseDto.IdEdicionTorneo)
	assert.Equal(t, edicionEquipoDto.IdEquipo, responseDto.IdEquipo)
}

/*func TestGetEdicionEquipos(t *testing.T) {

}*/

type MockEdicionEquipoService struct{}

// GetEdicionEquipos implements services.edicionEquipoServiceInterface.
func (m *MockEdicionEquipoService) GetEdicionEquipos(IdEdicionTorneo int) (dto.EdicionEquiposDto, errors.ApiError) {
	panic("unimplemented")
}

// InsertEdicionEquipos implements services.edicionEquipoServiceInterface.
func (m *MockEdicionEquipoService) InsertEdicionEquipos(edicionEquipoDto dto.EdicionEquipoDto) (dto.EdicionEquipoDto, errors.ApiError) {
	return edicionEquipoDto, nil
}

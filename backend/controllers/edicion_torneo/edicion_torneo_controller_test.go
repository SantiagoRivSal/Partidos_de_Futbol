package edicionTorneoController

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

func TestEdicionTorneoInsert(t *testing.T) {
	// Mock del servicio para simular la inserción de resultados
	services.EdicionTorneoService = &MockEdicionTorneoService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.POST("/torneo", EdicionTorneoInsert)

	// Crear un objeto EdicionTorneoDto simulado para la solicitud JSON
	edicionTorneoDto := dto.EdicionTorneoDto{
		Id: 1, IdTorneo: 1, Anio: 2024, SedeFinal: "Cordoba",
	}

	// Convertir el objeto EdicionTorneoDto en JSON
	payload, err := json.Marshal(edicionTorneoDto)
	assert.Nil(t, err)

	// Realizar una solicitud HTTP POST simulada a /torneo
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/torneo", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusCreated, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto EdicionTorneoDto
	var responseDto dto.EdicionTorneoDto
	err = json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos del torneo simulado
	assert.Equal(t, edicionTorneoDto.Id, responseDto.Id)
	assert.Equal(t, edicionTorneoDto.IdTorneo, responseDto.IdTorneo)
	assert.Equal(t, edicionTorneoDto.Anio, responseDto.Anio)
	assert.Equal(t, edicionTorneoDto.SedeFinal, responseDto.SedeFinal)

}

/*func TestGetEdicionTorneos(t *testing.T) {
	// Mock del servicio para simular la respuesta
	services.EdicionTorneoService = &MockEdicionTorneoService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.GET("/torneo/:id_torneo", GetEdicionTorneos)

	// ID de Torneo simulado
	idTorneo := 1

	// Realizar una solicitud HTTP GET simulada a /torneo/{id_torneo}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/torneo/"+strconv.Itoa(idTorneo), nil)
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto EdicionTorneosDto
	var responseDto dto.EdicionTorneosDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene los datos simulados
	assert.Equal(t, idTorneo, responseDto.IdTorneo)
	assert.Equal(t, 1, responseDto.Ediciones[0].Id)
	assert.Equal(t, "Torneo Simulado", responseDto.Ediciones[0].Nombre)
}*/

type MockEdicionTorneoService struct{}

// GetEdicionTorneos implements services.edicionTorneoServiceInterface.
func (m *MockEdicionTorneoService) GetEdicionTorneos(IdTorneo int) (dto.EdicionTorneosDto, errors.ApiError) {
	panic("unimplemented")
}

// GetEdicionTorneos implements services.edicionTorneoServiceInterface.
/*func (m *MockEdicionTorneoService) GetEdicionTorneos(IdTorneo int) (dto.EdicionTorneosDto, errors.ApiError) {
	// Simular la respuesta del servicio
	edicionTorneosDto := dto.EdicionTorneosDto{
		IdTorneo: IdTorneo,
		Ediciones: []dto.EdicionTorneoDto{
			{Id: 1, IdTorneo: 1, Anio: 2024, SedeFinal: "Cordoba"},
		},
	}
	return edicionTorneosDto, nil
}*/

// InsertEdicionTorneos implements services.edicionTorneoServiceInterface.
func (m *MockEdicionTorneoService) InsertEdicionTorneos(edicionTorneoDto dto.EdicionTorneoDto) (dto.EdicionTorneoDto, errors.ApiError) {
	return edicionTorneoDto, nil
}

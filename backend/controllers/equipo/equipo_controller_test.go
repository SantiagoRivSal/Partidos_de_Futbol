package equipoController

import (
	"backend/dto"
	"backend/services"
	"backend/utils/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockEquipoService struct{}

func (s *mockEquipoService) GetEquipos() (dto.EquiposDto, errors.ApiError) {
	return dto.EquiposDto{
		{Id: 1, Nombre: "Equipo1", Escudo: "Escudo1", IdPais: 100},
		{Id: 2, Nombre: "Equipo2", Escudo: "Escudo2", IdPais: 101},
	}, nil
}

func (s *mockEquipoService) GetEquipoById(id int) (dto.EquipoDto, errors.ApiError) {
	if id == 1 {
		return dto.EquipoDto{Id: 1, Nombre: "Equipo1", Escudo: "Escudo1", IdPais: 100}, nil
	}
	return dto.EquipoDto{}, errors.NewBadRequestApiError("equipo not found")
}

func (s *mockEquipoService) GetEquiposByIdPais(IdPais int) (dto.EquiposDto, errors.ApiError) {
	if IdPais == 100 {
		return dto.EquiposDto{
			{Id: 1, Nombre: "Equipo1", Escudo: "Escudo1", IdPais: 100},
		}, nil
	}
	return dto.EquiposDto{}, errors.NewBadRequestApiError("equipos no encontrados")
}

func TestGetEquipos(t *testing.T) {
	services.EquipoService = &mockEquipoService{}

	router := gin.Default()
	router.GET("/equipos", GetEquipos)

	req, _ := http.NewRequest("GET", "/equipos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `[{"id":1,"nombre":"Equipo1","escudo":"Escudo1","id_pais":100},{"id":2,"nombre":"Equipo2","escudo":"Escudo2","id_pais":101}]`
	assert.JSONEq(t, expected, w.Body.String())
}

func TestGetEquiposByIdPais(t *testing.T) {
	services.EquipoService = &mockEquipoService{}

	router := gin.Default()
	router.GET("/equipos/pais/:id_pais", GetEquiposByIdPais)

	req, _ := http.NewRequest("GET", "/equipos/pais/100", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `[{"id":1,"nombre":"Equipo1","escudo":"Escudo1","id_pais":100}]`
	assert.JSONEq(t, expected, w.Body.String())
}

func TestGetEquipoById(t *testing.T) {
	services.EquipoService = &mockEquipoService{}

	router := gin.Default()
	router.GET("/equipo/:id", GetEquipoById)

	req, _ := http.NewRequest("GET", "/equipo/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"id":1,"nombre":"Equipo1","escudo":"Escudo1","id_pais":100}`
	assert.JSONEq(t, expected, w.Body.String())

	req, _ = http.NewRequest("GET", "/equipo/2", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

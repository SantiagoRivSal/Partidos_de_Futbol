package equipoController

import (
	"backend/dto"
	service "backend/services"
	"backend/utils/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock del servicio
type MockEquipoService struct {
	mock.Mock
}

func (m *MockEquipoService) GetEquipos() (dto.EquiposDto, errors.ApiError) {
	args := m.Called()
	equipos := args.Get(0).(dto.EquiposDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Get(1).(errors.ApiError)
	}
	return equipos, apiErr
}

func (m *MockEquipoService) GetEquipoById(id int) (dto.EquipoDto, errors.ApiError) {
	args := m.Called(id)
	equipo := args.Get(0).(dto.EquipoDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Get(1).(errors.ApiError)
	}
	return equipo, apiErr
}

func (m *MockEquipoService) GetEquiposByIdPais(IdPais int) (dto.EquiposDto, errors.ApiError) {
	args := m.Called(IdPais)
	equipos := args.Get(0).(dto.EquiposDto)
	var apiErr errors.ApiError
	if args.Get(1) != nil {
		apiErr = args.Get(1).(errors.ApiError)
	}
	return equipos, apiErr
}

func TestGetEquipos(t *testing.T) {
	// Configuración de Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Creación del Mock del servicio
	mockService := new(MockEquipoService)
	service.EquipoService = mockService

	// Definición del comportamiento del Mock
	mockEquipos := dto.EquiposDto{
		{Id: 1, Nombre: "Equipo 1", Escudo: "escudo1.png", IdPais: 1},
		{Id: 2, Nombre: "Equipo 2", Escudo: "escudo2.png", IdPais: 2},
	}
	mockService.On("GetEquipos").Return(mockEquipos, nil)

	// Registro de la ruta
	router.GET("/equipos", GetEquipos)

	// Creación de la solicitud HTTP
	req, _ := http.NewRequest(http.MethodGet, "/equipos", nil)
	resp := httptest.NewRecorder()

	// Ejecución de la solicitud
	router.ServeHTTP(resp, req)

	// Aserciones
	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestGetEquiposByIdPais(t *testing.T) {
	// Configuración de Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Creación del Mock del servicio
	mockService := new(MockEquipoService)
	service.EquipoService = mockService

	// Definición del comportamiento del Mock
	mockEquipos := dto.EquiposDto{
		{Id: 1, Nombre: "Equipo 1", Escudo: "escudo1.png", IdPais: 1},
	}
	mockService.On("GetEquiposByIdPais", 1).Return(mockEquipos, nil)

	// Registro de la ruta
	router.GET("/equipos/:id_pais", GetEquiposByIdPais)

	// Creación de la solicitud HTTP
	req, _ := http.NewRequest(http.MethodGet, "/equipos/1", nil)
	resp := httptest.NewRecorder()

	// Ejecución de la solicitud
	router.ServeHTTP(resp, req)

	// Aserciones
	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestGetEquipoById(t *testing.T) {
	// Configuración de Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Creación del Mock del servicio
	mockService := new(MockEquipoService)
	service.EquipoService = mockService

	// Definición del comportamiento del Mock
	mockEquipo := dto.EquipoDto{Id: 1, Nombre: "Equipo 1", Escudo: "escudo1.png", IdPais: 1}
	mockService.On("GetEquipoById", 1).Return(mockEquipo, nil)

	// Registro de la ruta
	router.GET("/equipo/:id", GetEquipoById)

	// Creación de la solicitud HTTP
	req, _ := http.NewRequest(http.MethodGet, "/equipo/1", nil)
	resp := httptest.NewRecorder()

	// Ejecución de la solicitud
	router.ServeHTTP(resp, req)

	// Aserciones
	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

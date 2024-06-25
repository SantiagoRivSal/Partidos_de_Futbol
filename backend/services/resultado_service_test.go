package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockResultadoService struct {
	mock.Mock
}

func (m *MockResultadoService) GetResultados() (dto.ResultadosDto, e.ApiError) {
	args := m.Called()

	if args.Get(1) == nil {
		return args.Get(0).(dto.ResultadosDto), nil
	}

	return args.Get(0).(dto.ResultadosDto), args.Get(1).(e.ApiError)
}

func (m *MockResultadoService) GetResultadoByIdEdicionTorneo(IdEdicionTorneo int) (dto.ResultadoDto, e.ApiError) {
	args := m.Called(IdEdicionTorneo)

	if args.Get(1) == nil {
		return args.Get(0).(dto.ResultadoDto), nil
	}

	return args.Get(0).(dto.ResultadoDto), args.Get(1).(e.ApiError)
}

func (m *MockResultadoService) InsertResultados(resultadoDto dto.ResultadoDto) (dto.ResultadoDto, e.ApiError) {
	args := m.Called(resultadoDto)

	if args.Get(1) == nil {
		return args.Get(0).(dto.ResultadoDto), nil
	}

	return args.Get(0).(dto.ResultadoDto), args.Get(1).(e.ApiError)
}

func TestGetResultados(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockResultadoService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva resultados al llamar a GetResultados
	expectedResultadosDto := dto.ResultadosDto{
		{Id: 1, IdEdicionTorneo: 2023, Campeon: 10, Subcampeon: 20},
		{Id: 2, IdEdicionTorneo: 2024, Campeon: 30, Subcampeon: 40},
	}
	mockService.On("GetResultados").Return(expectedResultadosDto, nil)

	// Utilizar el mock en lugar del servicio real
	ResultadoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := ResultadoService.GetResultados()

	// Afirmar que la función devolvió los resultados esperados y no hay errores
	assert.Nil(t, err, "Error in GetResultados")
	assert.Equal(t, expectedResultadosDto, result)
}

func TestInsertResultados(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockResultadoService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva un resultados al llamar a InsertResultados
	expectedResultadoDto := dto.ResultadoDto{Id: 1, IdEdicionTorneo: 2023, Campeon: 10, Subcampeon: 20}
	mockService.On("InsertResultados", expectedResultadoDto).Return(expectedResultadoDto, nil)

	// Utilizar el mock en lugar del servicio real
	ResultadoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := ResultadoService.InsertResultados(expectedResultadoDto)

	// Afirmar que la función devolvió el resultados esperado y no hay errores
	assert.Nil(t, err, "Error in InsertResultados")
	assert.Equal(t, expectedResultadoDto, result)
}
func TestGetResultadoByIdEdicionTorneo(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockResultadoService)
	defer mockService.AssertExpectations(t)

	// Definir los datos esperados y el Id de prueba
	expectedResultadoDto := dto.ResultadoDto{Id: 1, IdEdicionTorneo: 2023, Campeon: 10, Subcampeon: 20}
	testIdEdicionTorneo := 2023

	// Configurar el mock para que devuelva el resultado esperado al llamar a GetResultadoByIdEdicionTorneo
	mockService.On("GetResultadoByIdEdicionTorneo", testIdEdicionTorneo).Return(expectedResultadoDto, nil)

	// Utilizar el mock en lugar del servicio real
	ResultadoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := ResultadoService.GetResultadoByIdEdicionTorneo(testIdEdicionTorneo)

	// Afirmar que la función devolvió el resultado esperado y no hay errores
	assert.Nil(t, err, "Error in GetResultadoByIdEdicionTorneo")
	assert.Equal(t, expectedResultadoDto, result)
}

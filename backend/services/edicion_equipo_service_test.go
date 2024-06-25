package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEdicionEquipoService struct {
	mock.Mock
}

func (m *MockEdicionEquipoService) GetEdicionEquipos(IdEdicionTorneo int) (dto.EdicionEquiposDto, e.ApiError) {
	args := m.Called(IdEdicionTorneo)

	if args.Get(1) == nil {
		return args.Get(0).(dto.EdicionEquiposDto), nil
	}

	return args.Get(0).(dto.EdicionEquiposDto), args.Get(1).(e.ApiError)
}

func (m *MockEdicionEquipoService) InsertEdicionEquipos(edicionEquipoDto dto.EdicionEquipoDto) (dto.EdicionEquipoDto, e.ApiError) {
	args := m.Called(edicionEquipoDto)

	if args.Get(1) == nil {
		return args.Get(0).(dto.EdicionEquipoDto), nil
	}

	return args.Get(0).(dto.EdicionEquipoDto), args.Get(1).(e.ApiError)
}

func TestInsertEdicionEquipos(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockEdicionEquipoService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva un resultados al llamar a InsertResultados
	expectedEdicionEquipoDto := dto.EdicionEquipoDto{Id: 1, IdEdicionTorneo: 1, IdEquipo: 1}
	mockService.On("InsertEdicionEquipos", expectedEdicionEquipoDto).Return(expectedEdicionEquipoDto, nil)

	// Utilizar el mock en lugar del servicio real
	EdicionEquipoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := EdicionEquipoService.InsertEdicionEquipos(expectedEdicionEquipoDto)

	// Afirmar que la función devolvió el resultados esperado y no hay errores
	assert.Nil(t, err, "Error in InsertEdicionEquipos")
	assert.Equal(t, expectedEdicionEquipoDto, result)
}

func TestGetEdicionEquipos(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockEdicionEquipoService)
	defer mockService.AssertExpectations(t)

	// Definir los datos esperados y el Id de prueba
	expectedEdicionEquiposDto := dto.EdicionEquiposDto{
		{Id: 1, IdEdicionTorneo: 1, IdEquipo: 1},
	}
	testIdEdicionTorneo := 1

	// Configurar el mock para que devuelva el resultado esperado al llamar a GetEdicionEquipos
	mockService.On("GetEdicionEquipos", testIdEdicionTorneo).Return(expectedEdicionEquiposDto, nil)

	// Utilizar el mock en lugar del servicio real
	EdicionEquipoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := EdicionEquipoService.GetEdicionEquipos(testIdEdicionTorneo)

	// Afirmar que la función devolvió el resultado esperado y no hay errores
	assert.Nil(t, err, "Error in GetEdicionEquipos")
	assert.Equal(t, expectedEdicionEquiposDto, result)
}

package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEdicionTorneoService struct {
	mock.Mock
}

func (m *MockEdicionTorneoService) GetEdicionTorneos(IdTorneo int) (dto.EdicionTorneosDto, e.ApiError) {
	args := m.Called(IdTorneo)

	if args.Get(1) == nil {
		return args.Get(0).(dto.EdicionTorneosDto), nil
	}

	return args.Get(0).(dto.EdicionTorneosDto), args.Get(1).(e.ApiError)
}

func (m *MockEdicionTorneoService) InsertEdicionTorneos(edicionTorneoDto dto.EdicionTorneoDto) (dto.EdicionTorneoDto, e.ApiError) {
	args := m.Called(edicionTorneoDto)

	if args.Get(1) == nil {
		return args.Get(0).(dto.EdicionTorneoDto), nil
	}

	return args.Get(0).(dto.EdicionTorneoDto), args.Get(1).(e.ApiError)
}

func TestInsertEdicionTorneos(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockEdicionTorneoService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva un resultados al llamar a InsertResultados
	expectedEdicionTorneoDto := dto.EdicionTorneoDto{Id: 1, IdTorneo: 1, Anio: 2024, SedeFinal: "Cordoba"}
	mockService.On("InsertEdicionTorneos", expectedEdicionTorneoDto).Return(expectedEdicionTorneoDto, nil)

	// Utilizar el mock en lugar del servicio real
	EdicionTorneoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := EdicionTorneoService.InsertEdicionTorneos(expectedEdicionTorneoDto)

	// Afirmar que la función devolvió el resultados esperado y no hay errores
	assert.Nil(t, err, "Error in InsertEdicionTorneos")
	assert.Equal(t, expectedEdicionTorneoDto, result)
}

func TestGetEdicionTorneos(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockEdicionTorneoService)
	defer mockService.AssertExpectations(t)

	// Definir los datos esperados y el Id de prueba
	expectedEdicionTorneosDto := dto.EdicionTorneosDto{
		{Id: 1, IdTorneo: 1, Anio: 2022, SedeFinal: "Cordoba"},
		{Id: 2, IdTorneo: 2, Anio: 2023, SedeFinal: "Rio de Janeiro"},
		{Id: 3, IdTorneo: 1, Anio: 2024, SedeFinal: "Lima"},
	}
	testIdTorneo := 1

	// Configurar el mock para que devuelva el resultado esperado al llamar a GetEdicionTorneos
	mockService.On("GetEdicionTorneos", testIdTorneo).Return(expectedEdicionTorneosDto, nil)

	// Utilizar el mock en lugar del servicio real
	EdicionTorneoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := EdicionTorneoService.GetEdicionTorneos(testIdTorneo)

	// Afirmar que la función devolvió el resultado esperado y no hay errores
	assert.Nil(t, err, "Error in GetEdicionTorneos")
	assert.Equal(t, expectedEdicionTorneosDto, result)
}

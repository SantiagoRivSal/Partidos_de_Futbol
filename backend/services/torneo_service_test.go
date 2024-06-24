package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTorneoService struct {
	mock.Mock
}

func (m *MockTorneoService) GetTorneos() (dto.TorneosDto, e.ApiError) {
	args := m.Called()

	if args.Get(1) == nil {
		return args.Get(0).(dto.TorneosDto), nil
	}

	return args.Get(0).(dto.TorneosDto), args.Get(1).(e.ApiError)
}

func TestGetTorneos(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockTorneoService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva las Paises al llamar a GetPaises
	expectedTorneosDto := dto.TorneosDto{
		{Id: 1, Nombre: "Torneo 1", Logo: "logo1.png", IdConfederacion: 1},
		{Id: 2, Nombre: "Torneo 2", Logo: "logo2.png", IdConfederacion: 1},
	}
	mockService.On("GetTorneos").Return(expectedTorneosDto, nil)

	// Utilizar el mock en lugar del servicio real
	TorneoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := TorneoService.GetTorneos()

	// Afirmar que la función devolvió los productos esperados y no hay errores
	assert.Nil(t, err, "Error in GetTorneos")
	assert.Equal(t, expectedTorneosDto, result)
}

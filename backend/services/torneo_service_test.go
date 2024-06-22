package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTorneoService es un mock para simular el Service torneoService
type MockTorneoService struct {
	mock.Mock
}

// GetTorneos simula la llamada a GetTorneos en el Service real
func (m *MockTorneoService) GetTorneos() (dto.TorneosDto, e.ApiError) {
	args := m.Called()
	if args.Get(1) == nil {
		return args.Get(0).(dto.TorneosDto), nil
	}

	return args.Get(0).(dto.TorneosDto), args.Get(1).(e.ApiError)
}

func TestGetTorneos(t *testing.T) {
	// Crear un mock del Service de torneos
	mockService := new(MockTorneoService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva torneos al llamar a GetTorneos
	expectedTorneosDto := dto.TorneosDto{
		{Id: 1, Nombre: "Torneo A", Logo: "logo1.jpg", IdConfederacion: 1},
		{Id: 2, Nombre: "Torneo B", Logo: "logo2.jpg", IdConfederacion: 1},
	}
	mockService.On("GetTorneos").Return(expectedTorneosDto)

	// Utilizar el mock en lugar del Service real
	TorneoService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := TorneoService.GetTorneos()

	// Afirmar que la función devolvió los torneos esperados y no hay errores
	assert.Nil(t, err, "Error in GetTorneos")
	assert.Equal(t, expectedTorneosDto, result)
}

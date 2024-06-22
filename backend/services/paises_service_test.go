package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPaisService es un mock para simular el Service paisService
type MockPaisService struct {
	mock.Mock
}

// GetPaises simula la llamada a GetPaises en el Service real
func (m *MockPaisService) GetPaises() (dto.PaisesDto, e.ApiError) {
	args := m.Called()

	if args.Get(1) == nil {
		return args.Get(0).(dto.PaisesDto), nil
	}

	return args.Get(0).(dto.PaisesDto), args.Get(1).(e.ApiError)
}

func TestGetPaises(t *testing.T) {
	// Crear un mock del Service de países
	mockService := new(MockPaisService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva países al llamar a GetPaises
	expectedPaisesDto := dto.PaisesDto{
		{Id: 1, Nombre: "País 1", Bandera: "bandera1.png", IdConfederacion: 1},
		{Id: 2, Nombre: "País 2", Bandera: "bandera2.png", IdConfederacion: 2},
	}
	mockService.On("GetPaises").Return(expectedPaisesDto)

	// Utilizar el mock en lugar del Service real
	PaisService = mockService

	// Llamar a la función que estamos probando con el mock
	paises, err := PaisService.GetPaises()

	// Afirmar que la función devolvió los países esperados y no hay errores
	assert.Nil(t, err, "Error in GetPaises")
	assert.Equal(t, expectedPaisesDto, paises)
}

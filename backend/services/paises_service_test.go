package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaisService struct {
	mock.Mock
}

func (m *MockPaisService) GetPaises() (dto.PaisesDto, e.ApiError) {
	args := m.Called()

	if args.Get(1) == nil {
		return args.Get(0).(dto.PaisesDto), nil
	}

	return args.Get(0).(dto.PaisesDto), args.Get(1).(e.ApiError)
}

func TestGetPaises(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockPaisService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva las Paises al llamar a GetPaises
	expectedPaisesDto := dto.PaisesDto{
		{Id: 1, Nombre: "Pais 1", Bandera: "bandera1.png"},
		{Id: 2, Nombre: "Pais 2", Bandera: "bandera2.png"},
	}
	mockService.On("GetPaises").Return(expectedPaisesDto, nil)

	// Utilizar el mock en lugar del servicio real
	PaisService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := PaisService.GetPaises()

	// Afirmar que la función devolvió los productos esperados y no hay errores
	assert.Nil(t, err, "Error in GetPaises")
	assert.Equal(t, expectedPaisesDto, result)
}

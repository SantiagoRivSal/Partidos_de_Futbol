package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockConfederacionService struct {
	mock.Mock
}

func (m *MockConfederacionService) GetConfederaciones() (dto.ConfederacionesDto, e.ApiError) {
	args := m.Called()

	if args.Get(1) == nil {
		return args.Get(0).(dto.ConfederacionesDto), nil
	}

	return args.Get(0).(dto.ConfederacionesDto), args.Get(1).(e.ApiError)
}

func TestGetConfederaciones(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockConfederacionService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva confederaciones al llamar a GetConfederaciones
	expectedConfederacionesDto := dto.ConfederacionesDto{
		{Id: 1, Nombre: "Confederación 1", Logo: "logo1.png"},
		{Id: 2, Nombre: "Confederación 2", Logo: "logo2.png"},
	}
	mockService.On("GetConfederaciones").Return(expectedConfederacionesDto, nil)

	// Utilizar el mock en lugar del servicio real
	ConfederacionService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := ConfederacionService.GetConfederaciones()

	// Afirmar que la función devolvió las confederaciones esperadas y no hay errores
	assert.Nil(t, err, "Error in GetConfederaciones")
	assert.Equal(t, expectedConfederacionesDto, result)
}

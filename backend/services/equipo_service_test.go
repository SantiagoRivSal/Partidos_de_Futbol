package services

import (
	"testing"

	"backend/dto"
	e "backend/utils/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEquipoService struct {
	mock.Mock
}

func (m *MockEquipoService) GetEquipos() (dto.EquiposDto, e.ApiError) {
	args := m.Called()

	if args.Get(1) == nil {
		return args.Get(0).(dto.EquiposDto), nil
	}

	return args.Get(0).(dto.EquiposDto), args.Get(1).(e.ApiError)
}

func (m *MockEquipoService) GetEquipoById(Id int) (dto.EquipoDto, e.ApiError) {
	args := m.Called(Id)

	if args.Get(1) == nil {
		return args.Get(0).(dto.EquipoDto), nil
	}

	return args.Get(0).(dto.EquipoDto), args.Get(1).(e.ApiError)
}

func (m *MockEquipoService) GetEquiposByIdPais(IdPais int) (dto.EquiposDto, e.ApiError) {
	args := m.Called(IdPais)

	if args.Get(1) == nil {
		return args.Get(0).(dto.EquiposDto), nil
	}

	return args.Get(0).(dto.EquiposDto), args.Get(1).(e.ApiError)
}

func TestGetEquipoById(t *testing.T) {
	mockService := new(MockEquipoService)
	defer mockService.AssertExpectations(t)

	expectedEquipoDto := dto.EquipoDto{Id: 1, Nombre: "Equipo 1", Escudo: "escudo1.png", IdPais: 1}
	testId := 1

	mockService.On("GetEquipoById", testId).Return(expectedEquipoDto, nil)

	EquipoService = mockService

	result, err := EquipoService.GetEquipoById(testId)
	assert.Nil(t, err, "Error in GetEquipoById")
	assert.Equal(t, expectedEquipoDto, result)
}

func TestGetEquipos(t *testing.T) {
	mockService := new(MockEquipoService)
	defer mockService.AssertExpectations(t)

	expectedEquiposDto := dto.EquiposDto{
		{Id: 1, Nombre: "Equipo 1", Escudo: "escudo1.png", IdPais: 1},
		{Id: 2, Nombre: "Equipo 2", Escudo: "escudo2.png", IdPais: 2},
	}
	mockService.On("GetEquipos").Return(expectedEquiposDto, nil)
	EquipoService = mockService

	result, err := EquipoService.GetEquipos()

	assert.Nil(t, err, "Error in GetEquipos")
	assert.Equal(t, expectedEquiposDto, result)
}

func TestGetEquiposByIdPais(t *testing.T) {
	mockService := new(MockEquipoService)
	defer mockService.AssertExpectations(t)

	idPais := 1
	expectedEquiposDto := dto.EquiposDto{
		{Id: 1, Nombre: "Equipo 1", Escudo: "escudo1.png", IdPais: 1},
	}

	mockService.On("GetEquiposByIdPais", idPais).Return(expectedEquiposDto, nil)

	EquipoService = mockService

	result, err := EquipoService.GetEquiposByIdPais(idPais)
	assert.Nil(t, err, "Error in GetEquiposByIdPais")
	assert.Equal(t, expectedEquiposDto, result)
}

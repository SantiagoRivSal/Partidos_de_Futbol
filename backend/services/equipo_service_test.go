package services

import (
	"testing"

	"backend/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEquipoCliente es un mock para simular el cliente equipoCliente
type MockEquipoCliente struct {
	mock.Mock
}

// GetEquipoById simula la llamada a GetEquipoById en el cliente real
func (m *MockEquipoCliente) GetEquipoById(id int) model.Equipo {
	args := m.Called(id)
	return args.Get(0).(model.Equipo)
}

// GetEquipos simula la llamada a GetEquipos en el cliente real
func (m *MockEquipoCliente) GetEquipos() model.Equipos {
	args := m.Called()
	return args.Get(0).(model.Equipos)
}

// GetEquiposByIdPais simula la llamada a GetEquiposByIdPais en el cliente real
func (m *MockEquipoCliente) GetEquiposByIdPais(IdPais int) model.Equipos {
	args := m.Called(IdPais)
	return args.Get(0).(model.Equipos)
}

func TestGetEquipoById(t *testing.T) {
	// Crear un mock del cliente de equipos
	mockCliente := new(MockEquipoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva un equipo al llamar a GetEquipoById
	expectedEquipo := model.Equipo{Id: 1, Nombre: "Equipo A", Escudo: "escudo1.jpg", IdPais: 1}
	mockCliente.On("GetEquipoById", 1).Return(expectedEquipo)

	// Utilizar el mock en lugar del cliente real
	equipoCliente = mockCliente

	// Llamar a la funcion que estamos probando con el mock
	service := &equipoService{}
	result, err := service.GetEquipoById(1)

	// Afirmar que la funcion devolvio el equipo esperado y no hay errores
	assert.Nil(t, err, "Error in GetEquipoById")
	assert.Equal(t, expectedEquipo.Nombre, result.Nombre)
	assert.Equal(t, expectedEquipo.Escudo, result.Escudo)
	assert.Equal(t, expectedEquipo.Id, result.Id)
	assert.Equal(t, expectedEquipo.IdPais, result.IdPais)
}

func TestGetEquipos(t *testing.T) {
	// Crear un mock del cliente de equipos
	mockCliente := new(MockEquipoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva equipos al llamar a GetEquipos
	expectedEquipos := model.Equipos{
		{Id: 1, Nombre: "Equipo A", Escudo: "escudo1.jpg", IdPais: 1},
		{Id: 2, Nombre: "Equipo B", Escudo: "escudo2.jpg", IdPais: 2},
	}
	mockCliente.On("GetEquipos").Return(expectedEquipos)

	// Utilizar el mock en lugar del cliente real
	equipoCliente = mockCliente

	// Llamar a la funcion que estamos probando con el mock
	service := &equipoService{}
	result, err := service.GetEquipos()

	// Afirmar que la funcion devolvio los equipos esperados y no hay errores
	assert.Nil(t, err, "Error in GetEquipos")
	assert.Equal(t, len(expectedEquipos), len(result))
}

func TestGetEquiposByIdPais(t *testing.T) {
	// Crear un mock del cliente de equipos
	mockCliente := new(MockEquipoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva equipos al llamar a GetEquiposByIdPais
	expectedEquipos := model.Equipos{
		{Id: 1, Nombre: "Equipo A", Escudo: "escudo1.jpg", IdPais: 1},
		{Id: 2, Nombre: "Equipo B", Escudo: "escudo2.jpg", IdPais: 1},
	}
	mockCliente.On("GetEquiposByIdPais", 1).Return(expectedEquipos)

	// Utilizar el mock en lugar del cliente real
	equipoCliente = mockCliente

	// Llamar a la funcion que estamos probando con el mock
	service := &equipoService{}
	result, err := service.GetEquiposByIdPais(1)

	// Afirmar que la funcion devolvio los equipos esperados y no hay errores
	assert.Nil(t, err, "Error in GetEquiposByIdPais")
	assert.Equal(t, len(expectedEquipos), len(result))
}

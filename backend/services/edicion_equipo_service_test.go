package services

import (
	"testing"

	"backend/dto"
	"backend/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEdicionEquipoCliente es un mock para simular el cliente edicionEquipoCliente
type MockEdicionEquipoCliente struct {
	mock.Mock
}

// GetEdicionEquipos simula la llamada a GetEdicionEquipos en el cliente real
func (m *MockEdicionEquipoCliente) GetEdicionEquipos(IdEdicionTorneo int) model.EdicionEquipos {
	args := m.Called(IdEdicionTorneo)
	return args.Get(0).(model.EdicionEquipos)
}

// InsertEdicionEquipos simula la llamada a InsertEdicionEquipos en el cliente real
func (m *MockEdicionEquipoCliente) InsertEdicionEquipos(edicionEquipo model.EdicionEquipo) model.EdicionEquipo {
	args := m.Called(edicionEquipo)
	return args.Get(0).(model.EdicionEquipo)
}

func TestGetEdicionEquipos(t *testing.T) {
	// Crear un mock del cliente de edicion de equipos
	mockCliente := new(MockEdicionEquipoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva edicion de equipos al llamar a GetEdicionEquipos
	expectedEdicionEquipos := model.EdicionEquipos{
		{Id: 1, IdEdicionTorneo: 1, IdEquipo: 101},
		{Id: 2, IdEdicionTorneo: 1, IdEquipo: 102},
	}
	mockCliente.On("GetEdicionEquipos", 1).Return(expectedEdicionEquipos)

	// Utilizar el mock en lugar del cliente real
	edicionEquipoCliente = mockCliente

	// Llamar a la funcion que estamos probando con el mock
	service := &edicionEquipoService{}
	result, err := service.GetEdicionEquipos(1)

	// Afirmar que la funcion devolvio las ediciones de equipos esperadas y no hay errores
	assert.Nil(t, err, "Error in GetEdicionEquipos")
	assert.Equal(t, len(expectedEdicionEquipos), len(result))
}

func TestInsertEdicionEquipos(t *testing.T) {
	// Crear un mock del cliente de edicion de equipos
	mockCliente := new(MockEdicionEquipoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva una edicion de equipos al llamar a InsertEdicionEquipos
	expectedEdicionEquipo := model.EdicionEquipo{Id: 1, IdEdicionTorneo: 1, IdEquipo: 101}
	mockCliente.On("InsertEdicionEquipos", mock.Anything).Return(expectedEdicionEquipo)

	// Utilizar el mock en lugar del cliente real
	edicionEquipoCliente = mockCliente

	// Llamar a la funcion que estamos probando con el mock
	service := &edicionEquipoService{}
	edicionEquipoDto := dto.EdicionEquipoDto{Id: 1, IdEdicionTorneo: 1, IdEquipo: 101}
	result, err := service.InsertEdicionEquipos(edicionEquipoDto)

	// Afirmar que la funcion devolvio la edicion de equipos esperada y no hay errores
	assert.Nil(t, err, "Error in InsertEdicionEquipos")
	assert.Equal(t, expectedEdicionEquipo.Id, result.Id)
	assert.Equal(t, expectedEdicionEquipo.IdEdicionTorneo, result.IdEdicionTorneo)
	assert.Equal(t, expectedEdicionEquipo.IdEquipo, result.IdEquipo)
}

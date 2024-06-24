/*package services

import (
	"testing"

	"backend/dto"
	"backend/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEdicionTorneoCliente es un mock para simular el cliente edicionTorneoCliente
type MockEdicionTorneoCliente struct {
	mock.Mock
}

// GetEdicionTorneos simula la llamada a GetEdicionTorneos en el cliente real
func (m *MockEdicionTorneoCliente) GetEdicionTorneos(IdTorneo int) model.EdicionTorneos {
	args := m.Called(IdTorneo)
	return args.Get(0).(model.EdicionTorneos)
}

// InsertEdicionTorneos simula la llamada a InsertEdicionTorneos en el cliente real
func (m *MockEdicionTorneoCliente) InsertEdicionTorneos(edicionTorneo model.EdicionTorneo) model.EdicionTorneo {
	args := m.Called(edicionTorneo)
	return args.Get(0).(model.EdicionTorneo)
}

func TestGetEdicionTorneos(t *testing.T) {
	// Crear un mock del cliente de edicion de torneos
	mockCliente := new(MockEdicionTorneoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva edicion de torneos al llamar a GetEdicionTorneos
	expectedEdicionTorneos := model.EdicionTorneos{
		{Id: 1, IdTorneo: 1, Anio: 2023, SedeFinal: "Estadio A"},
		{Id: 2, IdTorneo: 1, Anio: 2024, SedeFinal: "Estadio B"},
	}
	mockCliente.On("GetEdicionTorneos", 1).Return(expectedEdicionTorneos)

	// Utilizar el mock en lugar del cliente real
	edicionTorneoCliente = mockCliente

	// Llamar a la funcion que estamos probando con el mock
	service := &edicionTorneoService{}
	result, err := service.GetEdicionTorneos(1)

	// Afirmar que la funcion devolvio las ediciones de torneos esperadas y no hay errores
	assert.Nil(t, err, "Error in GetEdicionTorneos")
	assert.Equal(t, len(expectedEdicionTorneos), len(result))
}

func TestInsertEdicionTorneos(t *testing.T) {
	// Crear un mock del cliente de edicion de torneos
	mockCliente := new(MockEdicionTorneoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva una edicion de torneos al llamar a InsertEdicionTorneos
	expectedEdicionTorneo := model.EdicionTorneo{Id: 1, IdTorneo: 1, Anio: 2025, SedeFinal: "Estadio C"}
	mockCliente.On("InsertEdicionTorneos", mock.Anything).Return(expectedEdicionTorneo)

	// Utilizar el mock en lugar del cliente real
	edicionTorneoCliente = mockCliente

	// Llamar a la funcion que estamos probando con el mock
	service := &edicionTorneoService{}
	edicionTorneoDto := dto.EdicionTorneoDto{Id: 1, IdTorneo: 1, Anio: 2025, SedeFinal: "Estadio C"}
	result, err := service.InsertEdicionTorneos(edicionTorneoDto)

	// Afirmar que la funcion devolvio la edicion de torneos esperada y no hay errores
	assert.Nil(t, err, "Error in InsertEdicionTorneos")
	assert.Equal(t, expectedEdicionTorneo.Id, result.Id)
	assert.Equal(t, expectedEdicionTorneo.IdTorneo, result.IdTorneo)
	assert.Equal(t, expectedEdicionTorneo.Anio, result.Anio)
	assert.Equal(t, expectedEdicionTorneo.SedeFinal, result.SedeFinal)
}

package services

import (
	"testing"

	"backend/dto"
	"backend/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockResultadoCliente es un mock para simular el cliente resultadoCliente
type MockResultadoCliente struct {
	mock.Mock
}

// GetResultadoByIdEdicionTorneo simula la llamada a GetResultadoByIdEdicionTorneo en el cliente real
func (m *MockResultadoCliente) GetResultadoByIdEdicionTorneo(IdEdicionTorneo int) model.Resultado {
	args := m.Called(IdEdicionTorneo)
	return args.Get(0).(model.Resultado)
}

// GetResultados simula la llamada a GetResultados en el cliente real
func (m *MockResultadoCliente) GetResultados() model.Resultados {
	args := m.Called()
	return args.Get(0).(model.Resultados)
}

// InsertResultados simula la llamada a InsertResultados en el cliente real
func (m *MockResultadoCliente) InsertResultados(resultado model.Resultado) model.Resultado {
	args := m.Called(resultado)
	return args.Get(0).(model.Resultado)
}

func TestGetResultadoByIdEdicionTorneo(t *testing.T) {
	// Crear un mock del cliente de resultados
	mockCliente := new(MockResultadoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva un resultado al llamar a GetResultadoByIdEdicionTorneo
	expectedResultado := model.Resultado{
		Id:              1,
		IdEdicionTorneo: 1,
		Campeon:         2,
		Subcampeon:      3,
	}
	mockCliente.On("GetResultadoByIdEdicionTorneo", 1).Return(expectedResultado)

	// Utilizar el mock en lugar del cliente real
	resultadoCliente = mockCliente

	// Crear una instancia del servicio de resultados
	service := &resultadoService{}

	// Llamar a la función que estamos probando con el mock
	result, err := service.GetResultadoByIdEdicionTorneo(1)

	// Afirmar que la función devolvió el resultado esperado y no hay errores
	assert.Nil(t, err, "Error in GetResultadoByIdEdicionTorneo")
	assert.Equal(t, expectedResultado.Id, result.Id)
	assert.Equal(t, expectedResultado.IdEdicionTorneo, result.IdEdicionTorneo)
	assert.Equal(t, expectedResultado.Campeon, result.Campeon)
	assert.Equal(t, expectedResultado.Subcampeon, result.Subcampeon)
}

func TestGetResultados(t *testing.T) {
	// Crear un mock del cliente de resultados
	mockCliente := new(MockResultadoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva resultados al llamar a GetResultados
	expectedResultados := model.Resultados{
		{Id: 1, IdEdicionTorneo: 1, Campeon: 2, Subcampeon: 3},
		{Id: 2, IdEdicionTorneo: 2, Campeon: 5, Subcampeon: 8},
	}
	mockCliente.On("GetResultados").Return(expectedResultados)

	// Utilizar el mock en lugar del cliente real
	resultadoCliente = mockCliente

	// Crear una instancia del servicio de resultados
	service := &resultadoService{}

	// Llamar a la función que estamos probando con el mock
	resultados, err := service.GetResultados()

	// Afirmar que la función devolvió los resultados esperados y no hay errores
	assert.Nil(t, err, "Error in GetResultados")
	assert.Equal(t, len(expectedResultados), len(resultados))
}

func TestInsertResultados(t *testing.T) {
	// Crear un mock del cliente de resultados
	mockCliente := new(MockResultadoCliente)
	defer mockCliente.AssertExpectations(t)

	// Configurar el mock para que devuelva un resultado al llamar a InsertResultados
	expectedResultado := model.Resultado{
		Id:              1,
		IdEdicionTorneo: 1,
		Campeon:         3,
		Subcampeon:      6,
	}
	mockCliente.On("InsertResultados", mock.Anything).Return(expectedResultado)

	// Utilizar el mock en lugar del cliente real
	resultadoCliente = mockCliente

	// Crear una instancia del servicio de resultados
	service := &resultadoService{}

	// Llamar a la función que estamos probando con el mock
	resultadoDto := dto.ResultadoDto{
		IdEdicionTorneo: 1,
		Campeon:         3,
		Subcampeon:      6,
	}
	result, err := service.InsertResultados(resultadoDto)

	// Afirmar que la función devolvió el resultado esperado y no hay errores
	assert.Nil(t, err, "Error in InsertResultados")
	assert.Equal(t, expectedResultado.Id, result.Id)
	assert.Equal(t, expectedResultado.IdEdicionTorneo, result.IdEdicionTorneo)
	assert.Equal(t, expectedResultado.Campeon, result.Campeon)
	assert.Equal(t, expectedResultado.Subcampeon, result.Subcampeon)
}

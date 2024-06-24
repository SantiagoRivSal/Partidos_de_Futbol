package edicion_equipo

import (
	"testing"

	"backend/model"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDB es un mock de la base de datos
type MockDB struct {
	mock.Mock
}

// Find es un mock de la función Find de gorm.DB
func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

// Create es un mock de la función Create de gorm.DB
func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

// Where es un mock de la función Where de gorm.DB
func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	calledArgs := m.Called(append([]interface{}{query}, args...)...)
	return calledArgs.Get(0).(*gorm.DB)
}

// TestInsertEdicionEquipos prueba la función InsertEdicionEquipos
func TestInsertEdicionEquipos(t *testing.T) {
	// Crear un mock de la base de datos
	mockDB := new(MockDB)
	defer mockDB.AssertExpectations(t)

	// Configurar el mock para que devuelva una instancia de EdicionEquipo al llamar a Create
	expectedEdicionEquipo := model.EdicionEquipo{
		Id: 1, IdEdicionTorneo: 1, IdEquipo: 1,
	}
	mockDB.On("Create", &expectedEdicionEquipo).Return(&gorm.DB{})

	// Utilizar el mock en lugar de la conexión real a la base de datos
	Db = mockDB

	// Llamar a la función que estamos probando con el mock
	result := InsertEdicionEquipos(expectedEdicionEquipo)

	// Afirmar que la función devolvió el EdicionEquipo esperado
	assert.Equal(t, expectedEdicionEquipo, result)
}

// TestGetEdicionEquipos prueba la función GetEdicionEquipos
/*func TestGetEdicionEquipos(t *testing.T) {

		// Crear un mock de la base de datos
		mockDB := new(MockDB)
		defer mockDB.AssertExpectations(t)

		// Configurar el mock para que devuelva una instancia de EdicionEquipos al llamar a Where y Find
		expectedEdicionEquipos := model.EdicionEquipos{
			{Id: 1, IdEdicionTorneo: 1, IdEquipo: 1},
			{Id: 2, IdEdicionTorneo: 1, IdEquipo: 2},
		}

		// Configurar el mock para la llamada Where
		mockWhere := new(MockDB)
		mockFind := &gorm.DB{}

		mockDB.On("Where", "id_edicion_torneo = ?", 1).Return(mockWhere)
		mockWhere.On("Find", &expectedEdicionEquipos).Return(mockFind)

		// Utilizar el mock en lugar de la conexión real a la base de datos
		Db = mockDB

		// Llamar a la función que estamos probando con el mock
		result := GetEdicionEquipos(1)

		// Afirmar que la función devolvió los EdicionEquipos esperados
		assert.Equal(t, expectedEdicionEquipos, result)

}*/

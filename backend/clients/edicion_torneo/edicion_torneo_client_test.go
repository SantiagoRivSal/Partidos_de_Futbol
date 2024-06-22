package edicion_torneo

import (
	"testing"

	"backend/model"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	//"github.com/SantiagoRivSal/Partidos_de_Futbol/backend/model"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	callArgs := append([]interface{}{query}, args...)
	calledArgs := m.Called(callArgs...)
	return calledArgs.Get(0).(*gorm.DB)
}

func TestInsertEdicionTorneos(t *testing.T) {

	// Crear un mock de la base de datos
	mockDB := new(MockDB)
	defer mockDB.AssertExpectations(t)

	// Configurar el mock para que devuelva una instancia de EdicionTorneo al llamar a Create
	expectedEdicionTorneo := model.EdicionTorneo{
		Id: 1, IdTorneo: 1, Anio: 2024, SedeFinal: "Cordoba",
	}
	mockDB.On("Create", &expectedEdicionTorneo).Return(&gorm.DB{}, expectedEdicionTorneo)

	// Utilizar el mock en lugar de la conexión real a la base de datos
	Db = mockDB

	// Llamar a la función que estamos probando con el mock
	result := InsertEdicionTorneos(expectedEdicionTorneo)

	// Afirmar que la función devolvió el EdicionTorneoo esperado
	assert.Equal(t, expectedEdicionTorneo, result)

}

func TestGetEdicionTorneos(t *testing.T) {

}

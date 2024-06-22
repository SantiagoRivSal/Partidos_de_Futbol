package pais

import (
	"testing"

	"backend/model"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestGetPaises(t *testing.T) {
	// Crear un mock de la base de datos
	mockDB := new(MockDB)
	defer mockDB.AssertExpectations(t)

	// Configurar el mock para que devuelva directamente las Paises al llamar a Find
	expectedPaises := []model.Pais{
		{Id: 1, Nombre: "Pais 1", Bandera: "bandera1.png"},
		{Id: 2, Nombre: "Pais 2", Bandera: "bandera2.png"},
	}
	mockDB.On("Find", mock.AnythingOfType("*model.Paises"), mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*model.Paises)
		*dest = expectedPaises
	}).Return(&gorm.DB{})

	// Utilizar el mock en lugar de la conexión real a la base de datos
	Db = mockDB

	// Llamar a la función que estamos probando con el mock
	result := GetPaises()

	// Convertir model.Paises a []model.Confederacion para la comparación
	resultSlice := model.Paises(result)

	// Afirmar que la función devolvió los productos esperados
	assert.ElementsMatch(t, expectedPaises, resultSlice)
}

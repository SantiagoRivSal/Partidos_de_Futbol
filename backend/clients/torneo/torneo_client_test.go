package torneo

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

func TestGetTorneos(t *testing.T) {
	// Crear un mock de la base de datos
	mockDB := new(MockDB)
	defer mockDB.AssertExpectations(t)

	// Configurar el mock para que devuelva directamente las Torneos al llamar a Find
	expectedTorneos := []model.Torneo{
		{Id: 1, Nombre: "Torneo 1", Logo: "logo1.png", IdConfederacion: 1},
		{Id: 2, Nombre: "Torneo 2", Logo: "logo2.png", IdConfederacion: 1},
	}
	mockDB.On("Find", mock.AnythingOfType("*model.Torneos"), mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*model.Torneos)
		*dest = expectedTorneos
	}).Return(&gorm.DB{})

	// Utilizar el mock en lugar de la conexión real a la base de datos
	Db = mockDB

	// Llamar a la función que estamos probando con el mock
	result := GetTorneos()

	// Convertir model.Torneos a []model.Confederacion para la comparación
	resultSlice := model.Torneos(result)

	// Afirmar que la función devolvió los productos esperados
	assert.ElementsMatch(t, expectedTorneos, resultSlice)
}

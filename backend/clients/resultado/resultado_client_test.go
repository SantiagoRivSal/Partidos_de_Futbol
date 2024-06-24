package resultado

import (
	"backend/model"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDatabase es un mock de la interfaz Database
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDatabase) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	result := m.Called(query, args)
	return result.Get(0).(*gorm.DB)
}

func TestInsertResultados(t *testing.T) {
	mockDB := new(MockDatabase)
	Db = mockDB

	resultado := model.Resultado{
		Id:              1,
		IdEdicionTorneo: 2023,
		Campeon:         10,
		Subcampeon:      20,
	}

	// Crear una instancia de gorm.DB válida para el mock
	dbResult := &gorm.DB{}

	// Configurar el comportamiento esperado del mock
	mockDB.On("Create", &resultado).Return(dbResult)

	result := InsertResultados(resultado)

	// Comprobar que la función Create fue llamada con los argumentos correctos
	mockDB.AssertCalled(t, "Create", &resultado)
	// Comprobar que el resultado es el esperado
	assert.Equal(t, resultado, result)
}

func TestGetResultados(t *testing.T) {
	// Crear un mock de la base de datos
	mockDB := new(MockDatabase)
	defer mockDB.AssertExpectations(t)

	// Crear una instancia de gorm.DB válida para el mock
	dbResult := &gorm.DB{}

	// Configurar el mock para que devuelva directamente los resultados al llamar a Find
	expectedResultados := model.Resultados{
		{Id: 1, IdEdicionTorneo: 2023, Campeon: 10, Subcampeon: 20},
		{Id: 2, IdEdicionTorneo: 2024, Campeon: 30, Subcampeon: 40},
	}
	mockDB.On("Find", mock.AnythingOfType("*model.Resultados"), mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*model.Resultados)
		*dest = expectedResultados
	}).Return(dbResult)

	// Utilizar el mock en lugar de la conexión real a la base de datos
	Db = mockDB

	// Llamar a la función que estamos probando con el mock
	result := GetResultados()

	// Convertir model.Resultados a []model.Resultado para la comparación
	resultSlice := model.Resultados(result)

	// Afirmar que la función devolvió los resultados esperados
	assert.ElementsMatch(t, expectedResultados, resultSlice)
}

/*func TestGetResultadoByIdEdicionTorneo(t *testing.T) {
	mockDB := new(MockDatabase)
	Db = mockDB

	IdEdicionTorneo := 2023
	expectedResultado := model.Resultado{
		Id:              1,
		IdEdicionTorneo: IdEdicionTorneo,
		Campeon:         10,
		Subcampeon:      20,
	}

	// Crear una instancia de *gorm.DB que se puede devolver en las llamadas
	dbResult := &gorm.DB{}

	// Configurar el comportamiento esperado del mock
	mockDB.On("Where", "id_edicion_torneo = ?", []interface{}{IdEdicionTorneo}).Return(dbResult)
	mockDB.On("Find", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*model.Resultado)
		*dest = expectedResultado
	}).Return(dbResult)

	result := GetResultadoByIdEdicionTorneo(IdEdicionTorneo)

	// Comprobar que la función Where fue llamada con los argumentos correctos
	mockDB.AssertCalled(t, "Where", "id_edicion_torneo = ?", []interface{}{IdEdicionTorneo})
	// Comprobar que la función Find fue llamada con los argumentos correctos
	mockDB.AssertCalled(t, "Find", mock.Anything)
	// Comprobar que el resultado es el esperado
	assert.Equal(t, expectedResultado, result)
}*/

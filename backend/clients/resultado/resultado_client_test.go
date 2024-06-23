package resultado

import (
	"backend/model"
	"errors"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDatabase es una estructura que implementa la interfaz Database para ser usada en las pruebas.
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds...)
	db := &gorm.DB{}
	db.Error = args.Error(0)
	return db
}

func (m *MockDatabase) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	db := &gorm.DB{}
	db.Error = args.Error(0)
	return db
}

func (m *MockDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	mockArgs := m.Called(query, args...)
	db := &gorm.DB{}
	db.Error = mockArgs.Error(0)
	return db
}

func TestInsertResultados(t *testing.T) {
	mockDb := new(MockDatabase)
	Db = mockDb

	resultado := model.Resultado{
		Id:              1,
		IdEdicionTorneo: 100,
		Campeon:         1,
		Subcampeon:      2,
	}

	// Simulamos que la operación Create no produce errores
	mockDb.On("Create", &resultado).Return(nil)

	result := InsertResultados(resultado)

	assert.Equal(t, resultado, result)
	mockDb.AssertExpectations(t)
}

func TestGetResultados(t *testing.T) {
	mockDb := new(MockDatabase)
	Db = mockDb

	var resultados model.Resultados

	// Simulamos que la operación Find no produce errores
	mockDb.On("Find", &resultados).Return(nil)

	result := GetResultados()

	assert.Equal(t, resultados, result)
	mockDb.AssertExpectations(t)
}

func TestGetResultadoByIdEdicionTorneo(t *testing.T) {
	mockDb := new(MockDatabase)
	Db = mockDb

	idEdicionTorneo := 100
	var resultado model.Resultado

	// Simulamos que las operaciones Where y Find no producen errores
	mockDb.On("Where", "id_edicion_torneo = ?", idEdicionTorneo).Return(nil)
	mockDb.On("Find", &resultado).Return(nil)

	result := GetResultadoByIdEdicionTorneo(idEdicionTorneo)

	assert.Equal(t, resultado, result)
	mockDb.AssertExpectations(t)
}

func TestInsertResultadosWithError(t *testing.T) {
	mockDb := new(MockDatabase)
	Db = mockDb

	resultado := model.Resultado{
		Id:              1,
		IdEdicionTorneo: 100,
		Campeon:         1,
		Subcampeon:      2,
	}

	// Simulamos que la operación Create produce un error
	mockDb.On("Create", &resultado).Return(errors.New("some error"))

	result := InsertResultados(resultado)

	assert.Equal(t, resultado, result)
	mockDb.AssertExpectations(t)
}

func TestGetResultadoByIdEdicionTorneoWithError(t *testing.T) {
	mockDb := new(MockDatabase)
	Db = mockDb

	idEdicionTorneo := 100
	var resultado model.Resultado

	// Simulamos que las operaciones Where y Find producen errores
	mockDb.On("Where", "id_edicion_torneo = ?", idEdicionTorneo).Return(errors.New("some error"))
	mockDb.On("Find", &resultado).Return(errors.New("some error"))

	result := GetResultadoByIdEdicionTorneo(idEdicionTorneo)

	assert.Equal(t, resultado, result)
	mockDb.AssertExpectations(t)
}

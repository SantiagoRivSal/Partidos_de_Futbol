package equipo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type MockDatabase struct {
	db *gorm.DB
}

func (m *MockDatabase) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return m.db.Find(dest, conds...)
}

func (m *MockDatabase) Create(value interface{}) *gorm.DB {
	return m.db.Create(value)
}

func (m *MockDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	return m.db.Where(query, args...)
}

func TestGetEquipoById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	Db = &MockDatabase{gormDB}

	rows := sqlmock.NewRows([]string{"id", "nombre", "escudo", "id_pais"}).
		AddRow(1, "Equipo1", "Escudo1", 100)
	mock.ExpectQuery("^SELECT (.+) FROM `equipos` WHERE (.+)$").WillReturnRows(rows)

	equipo := GetEquipoById(1)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, 1, equipo.Id)
	assert.Equal(t, "Equipo1", equipo.Nombre)
	assert.Equal(t, "Escudo1", equipo.Escudo)
	assert.Equal(t, 100, equipo.IdPais)
}

func TestGetEquipos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	Db = &MockDatabase{gormDB}

	rows := sqlmock.NewRows([]string{"id", "nombre", "escudo", "id_pais"}).
		AddRow(1, "Equipo1", "Escudo1", 100).
		AddRow(2, "Equipo2", "Escudo2", 101)
	mock.ExpectQuery("^SELECT (.+) FROM `equipos`$").WillReturnRows(rows)

	equipos := GetEquipos()

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Len(t, equipos, 2)
	assert.Equal(t, 1, equipos[0].Id)
	assert.Equal(t, "Equipo1", equipos[0].Nombre)
	assert.Equal(t, 2, equipos[1].Id)
	assert.Equal(t, "Equipo2", equipos[1].Nombre)
}

func TestGetEquiposByIdPais(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	Db = &MockDatabase{gormDB}

	rows := sqlmock.NewRows([]string{"id", "nombre", "escudo", "id_pais"}).
		AddRow(1, "Equipo1", "Escudo1", 100).
		AddRow(2, "Equipo2", "Escudo2", 100)
	mock.ExpectQuery("^SELECT (.+) FROM `equipos` WHERE (.+)$").WillReturnRows(rows)

	equipos := GetEquiposByIdPais(100)

	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Len(t, equipos, 2)
	assert.Equal(t, 100, equipos[0].IdPais)
	assert.Equal(t, 100, equipos[1].IdPais)
}

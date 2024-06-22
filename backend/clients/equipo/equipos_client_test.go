package equipo

import (
	"backend/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetEquipoById(t *testing.T) {
	// Crear una conexión simulada a la base de datos
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Crear una instancia de GORM DB
	gormDB, err := gorm.Open("postgres", db)
	assert.NoError(t, err)
	Db = gormDB

	// Configurar la expectativa de la consulta SQL
	rows := sqlmock.NewRows([]string{"id", "nombre", "escudo", "id_pais"}).
		AddRow(1, "Equipo A", "escudo_a.png", 10)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "equipos" WHERE (id = $1)`)).
		WithArgs(1).
		WillReturnRows(rows)

	// Llamar a la función que se está probando
	result := GetEquipoById(1)

	// Verificar que las expectativas se cumplan
	assert.NoError(t, mock.ExpectationsWereMet())

	// Verificar el resultado
	expected := model.Equipo{Id: 1, Nombre: "Equipo A", Escudo: "escudo_a.png", IdPais: 10}
	assert.Equal(t, expected, result)
}

func TestGetEquipos(t *testing.T) {
	// Crear una conexión simulada a la base de datos
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Crear una instancia de GORM DB
	gormDB, err := gorm.Open("postgres", db)
	assert.NoError(t, err)
	Db = gormDB

	// Configurar la expectativa de la consulta SQL
	rows := sqlmock.NewRows([]string{"id", "nombre", "escudo", "id_pais"}).
		AddRow(1, "Equipo A", "escudo_a.png", 10).
		AddRow(2, "Equipo B", "escudo_b.png", 20)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "equipos"`)).
		WillReturnRows(rows)

	// Llamar a la función que se está probando
	result := GetEquipos()

	// Verificar que las expectativas se cumplan
	assert.NoError(t, mock.ExpectationsWereMet())

	// Verificar el resultado
	expected := model.Equipos{
		{Id: 1, Nombre: "Equipo A", Escudo: "escudo_a.png", IdPais: 10},
		{Id: 2, Nombre: "Equipo B", Escudo: "escudo_b.png", IdPais: 20},
	}
	assert.Equal(t, expected, result)
}

func TestGetEquiposByIdPais(t *testing.T) {
	// Crear una conexión simulada a la base de datos
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Crear una instancia de GORM DB
	gormDB, err := gorm.Open("postgres", db)
	assert.NoError(t, err)
	Db = gormDB

	// Configurar la expectativa de la consulta SQL
	rows := sqlmock.NewRows([]string{"id", "nombre", "escudo", "id_pais"}).
		AddRow(1, "Equipo A", "escudo_a.png", 10).
		AddRow(2, "Equipo C", "escudo_c.png", 10)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "equipos" WHERE (id_pais = $1)`)).
		WithArgs(10).
		WillReturnRows(rows)

	// Llamar a la función que se está probando
	result := GetEquiposByIdPais(10)

	// Verificar que las expectativas se cumplan
	assert.NoError(t, mock.ExpectationsWereMet())

	// Verificar el resultado
	expected := model.Equipos{
		{Id: 1, Nombre: "Equipo A", Escudo: "escudo_a.png", IdPais: 10},
		{Id: 2, Nombre: "Equipo C", Escudo: "escudo_c.png", IdPais: 10},
	}
	assert.Equal(t, expected, result)
}

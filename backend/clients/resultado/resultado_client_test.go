package resultado

import (
	"backend/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestInsertResultados(t *testing.T) {
	// Crear una conexión simulada a la base de datos
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Crear una instancia de GORM DB
	gormDB, err := gorm.Open("postgres", db)
	assert.NoError(t, err)
	Db = gormDB

	// Datos de prueba
	resultado := model.Resultado{
		Id:              1,
		IdEdicionTorneo: 100,
		Campeon:         1,
		Subcampeon:      2,
	}

	// Configurar la expectativa de la inserción SQL
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "resultados" ("id","id_edicion_torneo","campeon","subcampeon") VALUES ($1,$2,$3,$4) RETURNING "resultados"."id"`)).
		WithArgs(resultado.Id, resultado.IdEdicionTorneo, resultado.Campeon, resultado.Subcampeon).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Llamar a la función que se está probando
	result := InsertResultados(resultado)

	// Verificar que las expectativas se cumplan
	assert.NoError(t, mock.ExpectationsWereMet())

	// Verificar el resultado
	assert.Equal(t, resultado, result)
}

func TestGetResultados(t *testing.T) {
	// Crear una conexión simulada a la base de datos
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Crear una instancia de GORM DB
	gormDB, err := gorm.Open("postgres", db)
	assert.NoError(t, err)
	Db = gormDB

	// Configurar la expectativa de la consulta SQL
	rows := sqlmock.NewRows([]string{"id", "id_edicion_torneo", "campeon", "subcampeon"}).
		AddRow(1, 100, 1, 2).
		AddRow(2, 101, 3, 4)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resultados"`)).
		WillReturnRows(rows)

	// Llamar a la función que se está probando
	result := GetResultados()

	// Verificar que las expectativas se cumplan
	assert.NoError(t, mock.ExpectationsWereMet())

	// Verificar el resultado
	expected := model.Resultados{
		{Id: 1, IdEdicionTorneo: 100, Campeon: 1, Subcampeon: 2},
		{Id: 2, IdEdicionTorneo: 101, Campeon: 3, Subcampeon: 4},
	}
	assert.Equal(t, expected, result)
}

func TestGetResultadoByIdEdicionTorneo(t *testing.T) {
	// Crear una conexión simulada a la base de datos
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Crear una instancia de GORM DB
	gormDB, err := gorm.Open("postgres", db)
	assert.NoError(t, err)
	Db = gormDB

	// Configurar la expectativa de la consulta SQL
	rows := sqlmock.NewRows([]string{"id", "id_edicion_torneo", "campeon", "subcampeon"}).
		AddRow(1, 100, 1, 2)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resultados" WHERE (id_edicion_torneo = $1)`)).
		WithArgs(100).
		WillReturnRows(rows)

	// Llamar a la función que se está probando
	result := GetResultadoByIdEdicionTorneo(100)

	// Verificar que las expectativas se cumplan
	assert.NoError(t, mock.ExpectationsWereMet())

	// Verificar el resultado
	expected := model.Resultado{Id: 1, IdEdicionTorneo: 100, Campeon: 1, Subcampeon: 2}
	assert.Equal(t, expected, result)
}

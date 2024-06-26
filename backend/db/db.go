package db

import (
	confederacionClient "backend/clients/confederacion"
	edicionEquipoClient "backend/clients/edicion_equipos"
	edicionTorneoDetailClient "backend/clients/edicion_torneo"
	equipoClient "backend/clients/equipo"
	paisClient "backend/clients/pais"
	resultadoCliente "backend/clients/resultado"
	torneoClient "backend/clients/torneo"
	"backend/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")

	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	confederacionClient.Db = db
	resultadoCliente.Db = db
	edicionEquipoClient.Db = db
	edicionTorneoDetailClient.Db = db
	equipoClient.Db = db
	paisClient.Db = db
	torneoClient.Db = db
}

func StartDbEngine() {
	db.AutoMigrate(&model.Confederacion{})
	db.AutoMigrate(&model.EdicionEquipo{})
	db.AutoMigrate(&model.EdicionTorneo{})
	db.AutoMigrate(&model.Resultado{})
	db.AutoMigrate(&model.Equipo{})
	db.AutoMigrate(&model.Pais{})
	db.AutoMigrate(&model.Torneo{})
	log.Info("Finishing Migration Database Tables")
}

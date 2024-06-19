package db

import (
	confederacionClient "back/clients/confederacion"
	edicionEquipoClient "back/clients/edicion_equipos"
	edicionTorneoDetailClient "back/clients/edicion_torneo"
	equipoClient "back/clients/equipo"
	faseClient "back/clients/fase"
	paisClient "back/clients/pais"
	partidoClient "back/clients/partido"
	resultadoCliente "back/clients/resultado"
	torneoClient "back/clients/torneo"
	"back/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	DBName := "fifa"
	DBUser := "root"
	DBPass := "root"
	DBHost := "127.0.0.1"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True") //ver puerto

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
	faseClient.Db = db
	paisClient.Db = db
	partidoClient.Db = db
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

package db

import (
	//product "mvc-go/clients/product"

	confederacionClient "back/clients/confederacion"
	edicionEquipoClient "back/clients/edicion_equipo"
	edicionTorneoDetailClient "back/clients/edicion_torneo"
	equipoClient "back/clients/equipo"
	faseClient "back/clients/fase"
	paisClient "back/clients/pais"
	partidoClient "back/clients/partido"
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

	confederacion.Db = db 
	edicionEquipo.Db = db 
	edicionTorneoDetail.Db = db 
	equipo.Db = db 
	fase.Db = db 
	pais.Db = db 
	partido.Db = db 
	torneo.Db = db 
}

func StartDbEngine() {
	db.AutoMigrate(&model.Confederacion{})
	db.AutoMigrate(&model.EdicionEquipo{})
	db.AutoMigrate(&model.EdicionTorneo{})
	db.AutoMigrate(&model.Equipo{})
	db.AutoMigrate(&model.Fase{})
	db.AutoMigrate(&model.Pais{})
	db.AutoMigrate(&model.Partido{})
	db.AutoMigrate(&model.Torneo{})
	log.Info("Finishing Migration Database Tables")
}
package edicion_torneo

//ORM tradutcotr
import (
	"backend/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Database interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}

var Db Database

func InsertEdicionTorneos(edicionTorneo model.EdicionTorneo) model.EdicionTorneo {
	result := Db.Create(&edicionTorneo)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Nueva Edicion del Torneo Registrada: ", edicionTorneo.Id)
	return edicionTorneo
}

func GetEdicionTorneos(IdTorneo int) model.EdicionTorneos {
	var edicionTorneos model.EdicionTorneos

	Db.Where("id_torneo = ?", IdTorneo).Find(&edicionTorneos)
	log.Debug("Ediciones: ", edicionTorneos)

	return edicionTorneos
}

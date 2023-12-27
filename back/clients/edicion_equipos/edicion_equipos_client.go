package edicion_equipo

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertEdicionEquipos(edicionEquipo model.EdicionEquipo) model.EdicionEquipo {
	result := Db.Create(&edicionEquipo)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Equipos de la edicion Registrados: ", edicionEquipo.Id)
	return edicionEquipo
}

func GetEdicionEquipos(IdEdicionTorneo int) model.EdicionEquipos {
	var edicionEquipos model.EdicionEquipos
	Db.Where("id_edicion_torneo = ?", IdEdicionTorneo).Find(&edicionEquipos)

	log.Debug("Equipos Participantes de la edicion: ", edicionEquipos)

	return edicionEquipos
}

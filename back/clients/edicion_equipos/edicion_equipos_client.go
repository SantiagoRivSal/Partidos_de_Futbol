package edicion_equipos

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

func GetEdicionEquipos() model.EdicionEquipos {
	var edicionEquipos model.EdicionEquipos
	Db.Find(&edicionEquipos)

	log.Debug("Equipos Participantes de la edicion: ", edicionEquipos)

	return edicionEquipos
}
package edicion_torneo

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertEdicionTorneos(edicionTorneo model.EdicionTorneo) model.EdicionTorneo {
	result := Db.Create(&edicionTorneo)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Nueva Edicion del Torneo Registrada: ", edicionTorneo.Id)
	return edicionTorneo
}

func GetEdicionTorneos() model.EdicionTorneos {
	var edicionTorneos model.EdicionTorneos
	Db.Find(&edicionTorneos)

	log.Debug("Ediciones del Torneo: ", edicionTorneos)

	return edicionTorneos
}

/*func UpdateCampeon(campeon int, idEdicionTorneo int) int {
	result := Db.Model(&model.EdicionTorneo{}).Where("id = ?", idEdicionTorneo).Update("campeon", campeon)

	if result.Error != nil {
		log.Error("Equipo no participante")
	}
	return campeon
}*/

/*func UpdateSubcampeon(subcampeon int, idEdicionTorneo int) int {
	result := Db.Model(&model.EdicionTorneo{}).Where("id = ?", idEdicionTorneo).Update("subcampeon", subcampeon)

	if result.Error != nil {
		log.Error("Equipo no participante")
	}
	return subcampeon
}*/
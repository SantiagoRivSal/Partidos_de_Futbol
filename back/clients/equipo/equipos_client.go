package equipo

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetEquipos() model.Equipos {
	var equipos model.Equipos
	Db.Find(&equipos)

	log.Debug("Equipos: ", equipos)

	return equipos
}
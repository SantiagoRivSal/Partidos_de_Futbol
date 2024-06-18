package equipo

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetEquipoById(id int) model.Equipo {
	var equipo model.Equipo

	Db.Where("id = ?", id).Find(&equipo)
	log.Debug("Equipo: ", equipo)

	return equipo
}

func GetEquipos() model.Equipos {
	var equipos model.Equipos
	Db.Find(&equipos)

	log.Debug("Equipos: ", equipos)

	return equipos
}

func GetEquiposByIdPais(IdPais int) model.Equipos {
	var equipos model.Equipos

	Db.Where("id_pais = ?", IdPais).Find(&equipos)
	log.Debug("Equipos: ", equipos)

	return equipos
}

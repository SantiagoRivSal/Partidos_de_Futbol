package equipo

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

	Db.Where("id_pais = ?", 1).Find(&equipos)
	log.Debug("Equipos: ", equipos)

	return equipos
}

package pais

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetPaises() model.Paises {
	var paises model.Paises
	Db.Find(&paises)

	log.Debug("Paises: ", paises)

	return paises
}

func GetPaisesByIdConfederacion(IdConfederacion int) model.Paises {
	var paises model.Paises

	Db.Where("id_confederacion = ?", IdConfederacion).Find(&paises)
	log.Debug("Paises: ", paises)

	return paises
}

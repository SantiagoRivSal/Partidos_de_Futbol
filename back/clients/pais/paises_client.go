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
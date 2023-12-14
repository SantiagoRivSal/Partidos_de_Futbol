package confederacion

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetConfederaciones() model.Confederaciones {
	var confederaciones model.Confederaciones
	Db.Find(&confederaciones)

	log.Debug("Confederaciones: ", confederaciones)

	return confederaciones
}
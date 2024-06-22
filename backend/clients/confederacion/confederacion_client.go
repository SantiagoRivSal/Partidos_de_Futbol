package confederacion

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

func GetConfederaciones() model.Confederaciones {
	var confederaciones model.Confederaciones
	Db.Find(&confederaciones)

	log.Debug("Confederaciones: ", confederaciones)

	return confederaciones
}

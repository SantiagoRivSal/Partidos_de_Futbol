package torneo

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

func GetTorneos() model.Torneos {
	var torneos model.Torneos
	Db.Find(&torneos)

	log.Debug("Torneos: ", torneos)

	return torneos
}

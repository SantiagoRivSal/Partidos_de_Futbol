package torneo

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetTorneos() model.Torneos {
	var torneos model.Torneos
	Db.Find(&torneos)

	log.Debug("Torneos: ", torneos)

	return torneos
}

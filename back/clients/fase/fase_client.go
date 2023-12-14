package fase

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetFases() model.Fases {
	var fases model.Fases
	Db.Find(&fases)

	log.Debug("Fases: ", fases)

	return fases
}
package partido

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertPartidos(partido model.Partido) model.Partido {
	result := Db.Create(&partido)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Partido Registrado: ", partido.Id)
	return partido
}

func GetPartidos(IdEdicionTorneo, IdFase int) model.Partidos {
	var partidos model.Partidos
	Db.Where("id_edicion_torneo = ?", IdEdicionTorneo).Where("id_fase = ?", IdFase).Find(&partidos)

	log.Debug("Partidos: ", partidos)

	return partidos
}

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

func GetPartidos() model.Partidos {
	var partidos model.Partidos
	Db.Find(&partidos)

	log.Debug("Ediciones del Torneo: ", partidos)

	return partidos
}

/*func UpdateGolesLocal(golesLocal int, idPartido int) int {
	result := Db.Model(&model.Partido{}).Where("id = ?", idPartido).Update("goles_local", golesLocal)

	if result.Error != nil {
		log.Error("Equipo no participante")
	}
	return golesLocal
}*/

/*func UpdateGolesVisitante(golesVisitante int, idPartido int) int {
	result := Db.Model(&model.Partido{}).Where("id = ?", idPartido).Update("goles_visitante", golesVisitante)

	if result.Error != nil {
		log.Error("Equipo no participante")
	}
	return golesVisitante
}*/

/*func UpdateGanador(ganador int, idPartido int) int {
	result := Db.Model(&model.Partido{}).Where("id = ?", idPartido).Update("ganador", ganador)

	if result.Error != nil {
		log.Error("Equipo no participante")
	}
	return ganador
}*/
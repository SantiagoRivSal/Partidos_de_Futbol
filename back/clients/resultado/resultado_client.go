package resultado

//ORM tradutcotr
import (
	"back/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertResultados(resultado model.Resultado) model.Resultado {
	result := Db.Create(&resultado)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Nueva Resultado del Torneo Registrada: ", resultado.Id)
	return resultado
}

func GetResultados() model.Resultados {
	var resultados model.Resultados
	Db.Find(&resultados)

	log.Debug("Resultados de los Torneos: ", resultados)

	return resultados
}

func GetResultadoByIdEdicionTorneo(IdEdicionTorneo int) model.Resultado {
	var resultado model.Resultado

	Db.Where("id_edicion_torneo = ?", IdEdicionTorneo).Find(&resultado)
	log.Debug("Resultado: ", resultado)

	return resultado
}

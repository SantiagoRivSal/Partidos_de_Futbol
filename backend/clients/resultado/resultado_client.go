package resultado

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
	result := Db.Where("id_edicion_torneo = ?", IdEdicionTorneo).Find(&resultado)

	if result.Error != nil {
		log.Error("Error al obtener resultado por IdEdicionTorneo: ", result.Error)
	}
	log.Debug("Resultado: ", resultado)

	return resultado
}

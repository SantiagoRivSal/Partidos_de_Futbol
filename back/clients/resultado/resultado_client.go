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
	log.Debug("Nueva Edicion del Torneo Registrada: ", resultado.Id)
	return resultado
}

func GetResultados() model.Resultados {
	var resultados model.Resultados
	Db.Find(&resultados)

	log.Debug("Ediciones del Torneo: ", resultados)

	return resultados
}

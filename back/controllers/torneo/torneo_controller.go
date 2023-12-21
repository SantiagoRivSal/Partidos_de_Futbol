package torneoController

import (
	"back/dto"
	service "back/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func GetTorneos(c *gin.Context) {
	var torneosDto dto.TorneosDto
	torneosDto, err := service.TorneoService.GetTorneos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, torneosDto)
}

func GetTorneosByIdConfederacion(c *gin.Context) {
	log.Debug("Torneos id to load: " + c.Param("id_confederacion"))

	IdConfederacion, _ := strconv.Atoi(c.Param("id_confederacion"))
	var torneosDto dto.TorneosDto

	torneosDto, err := service.TorneoService.GetTorneosByIdConfederacion(IdConfederacion)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, torneosDto)
}

package edicionTorneoController

import (
	"back/dto"
	service "back/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func EdicionTorneoInsert(c *gin.Context) {
	var edicionTorneoDto dto.EdicionTorneoDto
	err := c.BindJSON(&edicionTorneoDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	edicionTorneoDto, er := service.EdicionTorneoService.InsertEdicionTorneo(edicionTorneoDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, edicionTorneoDto)
}

func GetEdicionTorneos(c *gin.Context) {
	var edicionTorneosDto dto.EdicionTorneosDto
	edicionTorneosDto, err := service.EdicionTorneoService.GetEdicionTorneos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, edicionTorneosDto)
}
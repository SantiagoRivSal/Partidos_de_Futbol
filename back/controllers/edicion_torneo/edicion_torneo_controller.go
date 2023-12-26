package edicionTorneoController

import (
	"back/dto"
	service "back/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func EdicionTorneoInsert(c *gin.Context) {
	var edicionTorneoDto dto.EdicionTorneoDto
	err := c.BindJSON(&edicionTorneoDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	edicionTorneoDto, er := service.EdicionTorneoService.InsertEdicionTorneos(edicionTorneoDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, edicionTorneoDto)
}

func GetEdicionTorneos(c *gin.Context) {
	log.Debug("Ediciones id to load: " + c.Param("id_torneo"))
	IdTorneo, _ := strconv.Atoi(c.Param("id_torneo"))
	var edicionTorneosDto dto.EdicionTorneosDto
	edicionTorneosDto, err := service.EdicionTorneoService.GetEdicionTorneos(IdTorneo)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, edicionTorneosDto)
}

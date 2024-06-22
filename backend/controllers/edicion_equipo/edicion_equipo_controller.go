package edicionEquipoController

import (
	"backend/dto"
	service "backend/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func EdicionEquipoInsert(c *gin.Context) {
	var edicionEquipoDto dto.EdicionEquipoDto
	err := c.BindJSON(&edicionEquipoDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	edicionEquipoDto, er := service.EdicionEquipoService.InsertEdicionEquipos(edicionEquipoDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, edicionEquipoDto)
}

func GetEdicionEquipos(c *gin.Context) {
	log.Debug("Ediciones id to load: " + c.Param("id_edicion_torneo"))
	IdEdicionTorneo, _ := strconv.Atoi(c.Param("id_edicion_torneo"))
	var edicionequiposDto dto.EdicionEquiposDto
	edicionequiposDto, err := service.EdicionEquipoService.GetEdicionEquipos(IdEdicionTorneo)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, edicionequiposDto)
}

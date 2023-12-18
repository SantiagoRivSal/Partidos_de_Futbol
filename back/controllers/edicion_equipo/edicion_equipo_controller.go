package edicionEquipoController

import (
	"back/dto"
	service "back/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func EdicionEquipoInsert(c *gin.Context) {
	var edicionEquipoDto dto.EdicionEquipoDto
	err := c.BindJSON(&edicionEquipoDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	edicionEquipoDto, er := service.EdicionEquipoService.InsertEdicionEquipo(edicionEquipoDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, edicionEquipoDto)
}

func GetEdicionEquipos(c *gin.Context) {
	var edicionequiposDto dto.EdicionEquiposDto
	edicionequiposDto, err := service.EdicionEquipoService.GetEdicionEquipos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, edicionequiposDto)
}


package partidoController

import (
	"back/dto"
	service "back/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func PartidosInsert(c *gin.Context) {
	var partidoDto dto.PartidoDto
	err := c.BindJSON(&partidoDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	partidoDto, er := service.PartidoService.InsertPartidos(partidoDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, partidoDto)
}

func GetPartidos(c *gin.Context) {
	var partidosDto dto.PartidosDto
	partidosDto, err := service.PartidoService.GetPartidos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, partidosDto)
}
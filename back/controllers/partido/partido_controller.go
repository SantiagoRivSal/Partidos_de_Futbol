package partidoController

import (
	"back/dto"
	service "back/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
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
	log.Debug("Torneo id to load: " + c.Param("id_edicion_torneo"))
	IdEdicionTorneo, _ := strconv.Atoi(c.Param("id_edicion_torneo"))

	// Agregar el nuevo parámetro
	log.Debug("Edicion id to load: " + c.Param("id_fase"))
	IdFase, _ := strconv.Atoi(c.Param("id_fase"))

	var partidosDto dto.PartidosDto
	partidosDto, er := service.PartidoService.GetPartidos(IdEdicionTorneo, IdFase)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, partidosDto)
}

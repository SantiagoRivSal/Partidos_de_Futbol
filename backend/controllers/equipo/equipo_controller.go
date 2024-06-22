package equipoController

import (
	"backend/dto"
	service "backend/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func GetEquipos(c *gin.Context) {
	var equiposDto dto.EquiposDto
	equiposDto, err := service.EquipoService.GetEquipos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, equiposDto)
}

func GetEquiposByIdPais(c *gin.Context) {
	log.Debug("Equipos id to load: " + c.Param("id_pais"))

	IdPais, _ := strconv.Atoi(c.Param("id_pais"))
	var equiposDto dto.EquiposDto

	equiposDto, err := service.EquipoService.GetEquiposByIdPais(IdPais)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, equiposDto)
}

func GetEquipoById(c *gin.Context) {
	log.Debug("Equipo id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var equipoDto dto.EquipoDto

	equipoDto, err := service.EquipoService.GetEquipoById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, equipoDto)
}

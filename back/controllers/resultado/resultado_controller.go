package resultadoController

import (
	"back/dto"
	service "back/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ResultadoInsert(c *gin.Context) {
	var resultadoDto dto.ResultadoDto
	err := c.BindJSON(&resultadoDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resultadoDto, er := service.ResultadoService.InsertResultados(resultadoDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, resultadoDto)
}

func GetResultados(c *gin.Context) {
	var resultadosDto dto.ResultadosDto
	resultadosDto, err := service.ResultadoService.GetResultados()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, resultadosDto)
}

func GetResultadoByIdEdicionTorneo(c *gin.Context) {
	log.Debug("Resultado id to load: " + c.Param("id_edicion_torneo"))

	id_edicion_torneo, _ := strconv.Atoi(c.Param("id_edicion_torneo"))
	var resultadoDto dto.ResultadoDto

	resultadoDto, err := service.ResultadoService.GetResultadoByIdEdicionTorneo(id_edicion_torneo)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, resultadoDto)
}

package resultadoController

import (
	"back/dto"
	service "back/services"
	"net/http"

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

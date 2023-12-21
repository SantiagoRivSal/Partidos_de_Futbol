package paisController

import (
	"back/dto"
	service "back/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func GetPaises(c *gin.Context) {
	var paisesDto dto.PaisesDto
	paisesDto, err := service.PaisService.GetPaises()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, paisesDto)
}

func GetPaisesByIdConfederacion(c *gin.Context) {
	log.Debug("Paises id to load: " + c.Param("id_confederacion"))

	IdConfederacion, _ := strconv.Atoi(c.Param("id_confederacion"))
	var paisesDto dto.PaisesDto

	paisesDto, err := service.PaisService.GetPaisesByIdConfederacion(IdConfederacion)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, paisesDto)
}

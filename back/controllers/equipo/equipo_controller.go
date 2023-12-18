package equipoController

import (
	"back/dto"
	service "back/services"
	"net/http"
	_ "strconv"

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
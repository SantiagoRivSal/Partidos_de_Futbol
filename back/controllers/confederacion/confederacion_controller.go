package confederacionController

import (
	"back/dto"
	service "back/services"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

func GetConfederaciones(c *gin.Context) {
	var confederacionesDto dto.ConfederacionesDto
	confederacionesDto, err := service.ConfederacionService.GetConfederaciones()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, confederacionesDto)
}
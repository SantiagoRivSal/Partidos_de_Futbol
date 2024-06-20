package paisController

import (
	"back/dto"
	service "back/services"
	"net/http"

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

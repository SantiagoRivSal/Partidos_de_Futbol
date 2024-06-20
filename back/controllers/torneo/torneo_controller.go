package torneoController

import (
	"back/dto"
	service "back/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTorneos(c *gin.Context) {
	var torneosDto dto.TorneosDto
	torneosDto, err := service.TorneoService.GetTorneos()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, torneosDto)
}

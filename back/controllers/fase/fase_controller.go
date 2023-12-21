package faseController

import (
	"back/dto"
	service "back/services"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

func GetFases(c *gin.Context) {
	var fasesDto dto.FasesDto
	fasesDto, err := service.FaseService.GetFases()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, fasesDto)
}

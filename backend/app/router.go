package app

import (
	"time"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Lista explícita de orígenes permitidos
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"}, // Añadir más encabezados si es necesario
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Solo funciona con orígenes específicos, no con '*'
		MaxAge:           12 * time.Hour,
	}))
}

func StartRoute() {
	mapUrls()

	log.Info("Starting server")
	router.Run(":4000")
}

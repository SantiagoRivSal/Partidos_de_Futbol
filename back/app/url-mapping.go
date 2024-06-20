package app

import (
	confederacionController "back/controllers/confederacion"
	edicionEquipoController "back/controllers/edicion_equipo"
	edicionTorneoController "back/controllers/edicion_torneo"
	equipoController "back/controllers/equipo"
	paisController "back/controllers/pais"
	resultadoController "back/controllers/resultado"
	torneoController "back/controllers/torneo"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Confederacion Mapping
	router.GET("/confederaciones", confederacionController.GetConfederaciones)

	//Equipos por Edicion Mapping
	router.GET("/equiposxedicion/:id_edicion_torneo", edicionEquipoController.GetEdicionEquipos)
	router.POST("/equipoxedicion", edicionEquipoController.EdicionEquipoInsert)

	//Edicion de Torneo Mapping
	router.GET("/ediciones_de_torneo/:id_torneo", edicionTorneoController.GetEdicionTorneos)
	router.POST("/edicion_de_torneo", edicionTorneoController.EdicionTorneoInsert)

	//Resultado de Torneo Mapping
	router.GET("/resultados", resultadoController.GetResultados)
	router.POST("/resultado", resultadoController.ResultadoInsert)
	router.GET("/resultadoxedicion/:id_edicion_torneos", resultadoController.GetResultadoByIdEdicionTorneo)

	// Equipos Mapping
	router.GET("/equipos", equipoController.GetEquipos)
	router.GET("/equiposxpais/:id_pais", equipoController.GetEquiposByIdPais)
	router.GET("/equipo/:id", equipoController.GetEquipoById)

	// Paises Mapping
	router.GET("/paises", paisController.GetPaises)

	// Torneos Mapping
	router.GET("/torneos", torneoController.GetTorneos)

	log.Info("Finishing mappings configurations")
}

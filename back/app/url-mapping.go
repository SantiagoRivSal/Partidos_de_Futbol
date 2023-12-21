package app

import (
	confederacionController "back/controllers/confederacion"
	edicionEquipoController "back/controllers/edicion_equipo"
	edicionTorneoController "back/controllers/edicion_torneo"
	equipoController "back/controllers/equipo"
	faseController "back/controllers/fase"
	paisController "back/controllers/pais"
	partidoController "back/controllers/partido"
	torneoController "back/controllers/torneo"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Confederacion Mapping
	router.GET("/confederaciones", confederacionController.GetConfederaciones)

	//Equipos por Edicion Mapping
	router.GET("/equiposxedicion", edicionEquipoController.GetEdicionEquipos)
	router.POST("/equipoxedicion", edicionEquipoController.EdicionEquipoInsert)

	//Edicion de Torneo Mapping
	router.GET("/ediciones_de_torneo", edicionTorneoController.GetEdicionTorneos)
	router.POST("/edicion_de_torneo", edicionTorneoController.EdicionTorneoInsert)

	// Equipos Mapping
	router.GET("/equipos", equipoController.GetEquipos)
	router.GET("/equiposxpais/:id_pais", equipoController.GetEquiposByIdPais)

	// Fases Mapping
	router.GET("/fases", faseController.GetFases)

	// Paises Mapping
	router.GET("/paises", paisController.GetPaises)
	router.GET("/paisesxconfederacion/:idConfederacion", paisController.GetPaisesByIdConfederacion)

	//PArtidos del Torneo Mapping
	router.GET("/partidos", partidoController.GetPartidos)
	router.POST("/partido", partidoController.PartidosInsert)

	// Torneos Mapping
	router.GET("/torneos", torneoController.GetTorneos)
	router.GET("/torneosxconfederacion/:idConfederacion", torneoController.GetTorneosByIdConfederacion)

	log.Info("Finishing mappings configurations")
}

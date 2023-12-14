package dto

type PartidoDto struct{
	Id int `json:"id"`
	IdEdicionTorneo int `json:"id_edicion_torneo"`
	IdFase int `json:"id_fase"`
	IdEquipoLocal int `json:"id_local"`
	IdEquipoVisitante int `json:"id_visitante"`
	GolesLocal int `json:"goles_local"`
	GolesVisitante int `json:"goles_visitante"`
	Ganador int `json:"ganador"`
}
type PartidosDto []PartidoDto

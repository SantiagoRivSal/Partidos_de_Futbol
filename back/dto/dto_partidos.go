package dto

type PartidoDto struct{
	Id int `json:"id"`
	IdEdicion int `json:"id_edicion"`
	IdFase int `json:"id_fase"`
	IdEquipoLocal int `json:"id_local"`
	IdEquipoVisitante int `json:"id_visitante"`
	GolesLocal int `json:"goles_local"`
	GolesVisitante int `json:"goles_visitante"`
	Ganador int `json:"ganador"`
}
type PartidosDto []PartidoDto
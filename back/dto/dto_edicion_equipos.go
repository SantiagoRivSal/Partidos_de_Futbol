package dto

type EdicionEquipoDto struct{
	Id int `json:"id"`
	IdEdicionTorneo int `json:"id_edicion_torneo"`
	IdEquipo int `json:"id_equipo"`
}
type EdicionEquiposDto []EdicionEquipoDto

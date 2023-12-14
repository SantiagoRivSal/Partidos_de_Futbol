package dto

type EdicionEquipoDto struct{
	Id int `json:"id"`
	IdEdicion int `json:"id_edicion"`
	IdEquipo int `json:"id_equipo"`
}
type EdicionEquiposDto []EdicionEquipoDto
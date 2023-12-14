package dto

type EquipoDto struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Escudo string `json:"escudo"`
	IdPais int `json:"id_pais"`
}
type PaisesDto []PaisDto
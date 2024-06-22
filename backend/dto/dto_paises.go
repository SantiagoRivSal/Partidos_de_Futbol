package dto

type PaisDto struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Bandera string `json:"bandera"`
	IdConfederacion int `json:"id_confederacion"`
}
type PaisesDto []PaisDto
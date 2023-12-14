package dto

type FaseDto struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
}
type FasesDto []FaseDto
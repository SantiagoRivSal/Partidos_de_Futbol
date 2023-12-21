package dto

type FaseDto struct {
	Id              int    `json:"id"`
	Nombre          string `json:"nombre"`
	CantidadEquipos int    `json:"cantidad_equipos"`
}
type FasesDto []FaseDto

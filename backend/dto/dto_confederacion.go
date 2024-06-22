package dto

type ConfederacionDto struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Logo string `json:"logo"`
}
type ConfederacionesDto []ConfederacionDto
package dto

type TorneoDto struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Logo string `json:"logo"`
	IdConfederacion int `json:"id_confederacion"`

}
type TorneosDto []TorneoDto
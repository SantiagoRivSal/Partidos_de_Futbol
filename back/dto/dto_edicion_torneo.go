package dto

type EdicionTorneoDto struct{
	Id int `json:"id"`
	IdTorneo int `json:"id_torneo"`
	Anio int `json:"anio"`
	Campeon int `json:"campeon"`
	Subcampeon int `json:"subcampeon"`
}
type TorneosDto []TorneoDto
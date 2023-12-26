package dto

type ResultadoDto struct {
	Id              int `json:"id"`
	IdEdicionTorneo int `json:"id_edicion_torneo"`
	Campeon         int `json:"campeon"`
	Subcampeon      int `json:"subcampeon"`
}
type ResultadosDto []ResultadoDto

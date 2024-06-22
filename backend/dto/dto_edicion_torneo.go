package dto

type EdicionTorneoDto struct {
	Id        int    `json:"id"`
	IdTorneo  int    `json:"id_torneo"`
	Anio      int    `json:"anio"`
	SedeFinal string `json:"sede_final"`
}
type EdicionTorneosDto []EdicionTorneoDto

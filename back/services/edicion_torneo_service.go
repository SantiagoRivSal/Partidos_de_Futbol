package services

import (
	edicionTorneoCliente "back/clients/edicion_torneo"
	"back/dto"
	"back/model"
	e "back/utils/errors"
)

type edicionTorneoService struct{}

type edicionTorneoServiceServiceInterface interface {
	GetEdicionTorneos() (dto.EdicionTorneosDto, e.ApiError)
	InsertEdicionTorneos(edicionTorneoDto dto.EdicionTorneoDto) (dto.EdicionTorneoDto, e.ApiError)
}

var (
	EdicionTorneoService edicionTorneoServiceInterface
)

func init() {
	EdicionTorneoService = &edicionTorneoService{}
}

func (s *edicionTorneoService) GetEdicionTorneos() (dto.EdicionTorneosDto, e.ApiError) {

	var edicionTorneos model.EdicionTorneos = edicionTorneoCliente.GetEdicionTorneos()
	var edicionTorneosDto dto.EdicionTorneosDto
	if len(edicionTorneos) == 0 {
		return edicionTorneosDto, e.NewBadRequestApiError("Ediciones de torneos no encontradas")
	}
	for _, edicionTorneo := range edicionTorneos {
		var edicionTorneoDto dto.EdicionTorneoDto
		edicionTorneoDto.IdTorneo = edicionTorneo.IdTorneo
		edicionTorneoDto.Anio = edicionTorneo.Anio
		edicionTorneoDto.Campeon = edicionTorneo.Campeon
		edicionTorneoDto.Anio = edicionTorneo.Anio
		edicionTorneoDto.Subcampeon = edicionTorneo.Subcampeon

		edicionTorneosDto = append(edicionTorneosDto, edicionTorneoDto)
	}

	return edicionTorneosDto, nil
}

func (s *edicionTorneoService) InsertEdicionTorneos(edicionTorneoDto dto.EdicionTorneoDto) (dto.EdicionTorneoDto, e.ApiError) {

	var edicionTorneo model.EdicionTorneo

	edicionTorneo.IdTorneo= edicionTorneoDto.IdTorneo
	edicionTorneo.Anio = edicionTorneoDto.Anio
	edicionTorneo.Campeon = edicionTorneoDto.Campeon
	edicionTorneo.Anio = edicionTorneoDto.Anio
	edicionTorneo.Subcampeon = edicionTorneoDto.Subcampeon

	edicionTorneo = edicionTorneoCliente.InsertEdicionTorneos(edicionTorneo)

	edicionTorneoDto.IdTorneo = edicionTorneo.IdTorneo
	edicionTorneoDto.Anio = edicionTorneo.Anio
	edicionTorneoDto.Campeon = edicionTorneo.Campeon
	edicionTorneoDto.Anio = edicionTorneo.Anio
	edicionTorneoDto.Subcampeon = edicionTorneo.Subcampeon

	return edicionTorneoDto, nil

}
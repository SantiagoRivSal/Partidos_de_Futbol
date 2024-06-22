package services

import (
	edicionTorneoCliente "backend/clients/edicion_torneo"
	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type edicionTorneoService struct{}

type edicionTorneoServiceInterface interface {
	GetEdicionTorneos(IdTorneo int) (dto.EdicionTorneosDto, e.ApiError)
	InsertEdicionTorneos(edicionTorneoDto dto.EdicionTorneoDto) (dto.EdicionTorneoDto, e.ApiError)
}

var (
	EdicionTorneoService edicionTorneoServiceInterface
)

func init() {
	EdicionTorneoService = &edicionTorneoService{}
}

func (s *edicionTorneoService) GetEdicionTorneos(IdTorneo int) (dto.EdicionTorneosDto, e.ApiError) {

	var edicionTorneos model.EdicionTorneos = edicionTorneoCliente.GetEdicionTorneos(IdTorneo)
	var edicionTorneosDto dto.EdicionTorneosDto
	if len(edicionTorneos) == 0 {
		return edicionTorneosDto, e.NewBadRequestApiError("Ediciones de torneos no encontradas")
	}
	for _, edicionTorneo := range edicionTorneos {
		var edicionTorneoDto dto.EdicionTorneoDto
		edicionTorneoDto.IdTorneo = edicionTorneo.IdTorneo
		edicionTorneoDto.Id = edicionTorneo.Id
		edicionTorneoDto.Anio = edicionTorneo.Anio
		edicionTorneoDto.SedeFinal = edicionTorneo.SedeFinal

		edicionTorneosDto = append(edicionTorneosDto, edicionTorneoDto)
	}

	return edicionTorneosDto, nil
}

func (s *edicionTorneoService) InsertEdicionTorneos(edicionTorneoDto dto.EdicionTorneoDto) (dto.EdicionTorneoDto, e.ApiError) {

	var edicionTorneo model.EdicionTorneo

	edicionTorneo.IdTorneo = edicionTorneoDto.IdTorneo
	edicionTorneo.Anio = edicionTorneoDto.Anio
	edicionTorneo.Id = edicionTorneoDto.Id
	edicionTorneo.SedeFinal = edicionTorneoDto.SedeFinal

	edicionTorneo = edicionTorneoCliente.InsertEdicionTorneos(edicionTorneo)

	edicionTorneoDto.IdTorneo = edicionTorneo.IdTorneo
	edicionTorneoDto.Anio = edicionTorneo.Anio
	edicionTorneoDto.Id = edicionTorneo.Id
	edicionTorneoDto.SedeFinal = edicionTorneo.SedeFinal

	return edicionTorneoDto, nil

}

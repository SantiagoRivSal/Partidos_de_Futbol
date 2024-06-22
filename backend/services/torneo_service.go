package services

import (
	torneoCliente "backend/clients/torneo"
	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type torneoService struct{}

type torneoServiceInterface interface {
	GetTorneos() (dto.TorneosDto, e.ApiError)
}

var (
	TorneoService torneoServiceInterface
)

func init() {
	TorneoService = &torneoService{}
}

func (s *torneoService) GetTorneos() (dto.TorneosDto, e.ApiError) {

	var torneos model.Torneos = torneoCliente.GetTorneos()
	var torneosDto dto.TorneosDto
	if len(torneos) == 0 {
		return torneosDto, e.NewBadRequestApiError("torneos no encontrados")
	}
	for _, torneo := range torneos {
		var torneoDto dto.TorneoDto
		torneoDto.Nombre = torneo.Nombre
		torneoDto.Logo = torneo.Logo
		torneoDto.Id = torneo.Id
		torneoDto.IdConfederacion = torneo.IdConfederacion

		torneosDto = append(torneosDto, torneoDto)
	}

	return torneosDto, nil
}

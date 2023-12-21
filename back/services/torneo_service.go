package services

import (
	torneoCliente "back/clients/torneo"
	"back/dto"
	"back/model"
	e "back/utils/errors"
)

type torneoService struct{}

type torneoServiceInterface interface {
	GetTorneos() (dto.TorneosDto, e.ApiError)
	GetTorneosByIdConfederacion(IdConfederacion int) (dto.TorneosDto, e.ApiError)
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

func (s *torneoService) GetTorneosByIdConfederacion(IdConfederacion int) (dto.TorneosDto, e.ApiError) {

	var torneos model.Torneos = torneoCliente.GetTorneosByIdConfederacion(IdConfederacion)
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

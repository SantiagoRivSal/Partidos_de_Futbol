package services

import (
	equipoCliente "back/clients/equipo"
	"back/dto"
	"back/model"
	e "back/utils/errors"
)

type equipoService struct{}

type equipoServiceInterface interface {
	GetEquipos() (dto.EquiposDto, e.ApiError)
	GetEquiposByIdPais(IdPais int) (dto.EquiposDto, e.ApiError)
}

var (
	EquipoService equipoServiceInterface
)

func init() {
	EquipoService = &equipoService{}
}

func (s *equipoService) GetEquipos() (dto.EquiposDto, e.ApiError) {

	var equipos model.Equipos = equipoCliente.GetEquipos()
	var equiposDto dto.EquiposDto
	if len(equipos) == 0 {
		return equiposDto, e.NewBadRequestApiError("equipos no encontrados")
	}
	for _, equipo := range equipos {
		var equipoDto dto.EquipoDto
		equipoDto.Nombre = equipo.Nombre
		equipoDto.Escudo = equipo.Escudo
		equipoDto.Id = equipo.Id
		equipoDto.IdPais = equipo.IdPais

		equiposDto = append(equiposDto, equipoDto)
	}

	return equiposDto, nil
}

func (s *equipoService) GetEquiposByIdPais(IdPais int) (dto.EquiposDto, e.ApiError) {

	var equipos model.Equipos = equipoCliente.GetEquiposByIdPais(IdPais)
	var equiposDto dto.EquiposDto
	if len(equipos) == 0 {
		return equiposDto, e.NewBadRequestApiError("equipos no encontrados")
	}
	for _, equipo := range equipos {
		var equipoDto dto.EquipoDto
		equipoDto.Nombre = equipo.Nombre
		equipoDto.Escudo = equipo.Escudo
		equipoDto.Id = equipo.Id
		equipoDto.IdPais = equipo.IdPais

		equiposDto = append(equiposDto, equipoDto)
	}

	return equiposDto, nil
}

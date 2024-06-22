package services

import (
	equipoCliente "backend/clients/equipo"
	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type equipoService struct{}

type equipoServiceInterface interface {
	GetEquipos() (dto.EquiposDto, e.ApiError)
	GetEquipoById(id int) (dto.EquipoDto, e.ApiError)
	GetEquiposByIdPais(IdPais int) (dto.EquiposDto, e.ApiError)
}

var (
	EquipoService equipoServiceInterface
)

func init() {
	EquipoService = &equipoService{}
}

func (s *equipoService) GetEquipoById(id int) (dto.EquipoDto, e.ApiError) {

	var equipo model.Equipo = equipoCliente.GetEquipoById(id)
	var equipoDto dto.EquipoDto

	if equipo.Id == 0 {
		return equipoDto, e.NewBadRequestApiError("equipo not found")
	}
	equipoDto.Nombre = equipo.Nombre
	equipoDto.Escudo = equipo.Escudo
	equipoDto.Id = equipo.Id
	equipoDto.IdPais = equipo.IdPais
	return equipoDto, nil
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

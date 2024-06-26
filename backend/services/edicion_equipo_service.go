package services

import (
	edicionEquipoCliente "backend/clients/edicion_equipos"
	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type edicionEquipoService struct{}

type edicionEquipoServiceInterface interface {
	GetEdicionEquipos(IdEdicionTorneo int) (dto.EdicionEquiposDto, e.ApiError)
	InsertEdicionEquipos(edicionEquipoDto dto.EdicionEquipoDto) (dto.EdicionEquipoDto, e.ApiError)
}

var (
	EdicionEquipoService edicionEquipoServiceInterface
)

func init() {
	EdicionEquipoService = &edicionEquipoService{}
}

func (s *edicionEquipoService) GetEdicionEquipos(IdEdicionTorneo int) (dto.EdicionEquiposDto, e.ApiError) {

	var edicionEquipos model.EdicionEquipos = edicionEquipoCliente.GetEdicionEquipos(IdEdicionTorneo)
	var edicionEquiposDto dto.EdicionEquiposDto
	if len(edicionEquipos) == 0 {
		return edicionEquiposDto, e.NewBadRequestApiError("equipos de ediciones no encontradas")
	}
	for _, edicionEquipo := range edicionEquipos {
		var edicionEquipoDto dto.EdicionEquipoDto
		edicionEquipoDto.IdEdicionTorneo = edicionEquipo.IdEdicionTorneo
		edicionEquipoDto.IdEquipo = edicionEquipo.IdEquipo
		edicionEquipoDto.Id = edicionEquipo.Id

		edicionEquiposDto = append(edicionEquiposDto, edicionEquipoDto)
	}

	return edicionEquiposDto, nil
}

func (s *edicionEquipoService) InsertEdicionEquipos(edicionEquipoDto dto.EdicionEquipoDto) (dto.EdicionEquipoDto, e.ApiError) {

	var edicionEquipo model.EdicionEquipo

	edicionEquipo.IdEdicionTorneo = edicionEquipoDto.IdEdicionTorneo
	edicionEquipo.IdEquipo = edicionEquipoDto.IdEquipo
	edicionEquipo.Id = edicionEquipoDto.Id

	edicionEquipo = edicionEquipoCliente.InsertEdicionEquipos(edicionEquipo)

	edicionEquipoDto.IdEdicionTorneo = edicionEquipo.IdEdicionTorneo
	edicionEquipoDto.IdEquipo = edicionEquipo.IdEquipo
	edicionEquipoDto.Id = edicionEquipo.Id

	return edicionEquipoDto, nil

}

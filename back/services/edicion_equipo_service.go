package services

import (
	edicionEquipoCliente "back/clients/edicion_equipos"
	"back/dto"
	"back/model"
	e "back/utils/errors"
)

type edicionEquipoService struct{}

type edicionEquipoServiceInterface interface {
	GetEdicionEquipos() (dto.EdicionEquiposDto, e.ApiError)
	InsertEdicionEquipos(edicionEquipoDto dto.EdicionEquipoDto) (dto.EdicionEquipoDto, e.ApiError)
}

var (
	EdicionEquipoService edicionEquipoServiceInterface
)

func init() {
	EdicionEquipoService = &edicionEquipoService{}
}

func (s *edicionEquipoService) GetEdicionEquipos() (dto.EdicionEquiposDto, e.ApiError) {

	var edicionEquipos model.EdicionEquipos = edicionEquipoCliente.GetEdicionEquipos()
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

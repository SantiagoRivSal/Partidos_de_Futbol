package services

import (
	edicionEquipoCliente "back/clients/edicion_equipos"
	"back/dto"
	"back/model"
	e "back/utils/errors"
)

type edicionEquipoService struct{}

type edicionEquipoServiceServiceInterface interface {
	GetEdicionEquipos() (dto.EdicionEquipoesDto, e.ApiError)
	InsertEdicionEquipos(edicionEquipoDto dto.EdicionEquipoDto) (dto.EdicionEquipoDto, e.ApiError)
}

var (
	EdicionEquipoService edicionEquipoServiceServiceInterface
)

func init() {
	EdicionEquipoService = &edicionEquipoService{}
}

func (s *edicionEquipoService) GetEdicionEquipos() (dto.EdicionEquiposDto, e.ApiError) {

	var edicionEquipos model.EdicionEquipoes = edicionEquipoCliente.GetEdicionEquipos()
	var edicionEquiposDto dto.EdicionEquipoesDto
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

	return edicionEquipoesDto, nil
}

func (s *edicionEquipoService) InsertEdicionEquipos(edicionEquipoDto dto.EdicionEquipoDto) (dto.EdicionEquipoDto, e.ApiError) {

}

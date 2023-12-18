package services

import (
	partidoCliente "back/clients/partido"
	"back/dto"
	"back/model"
	e "back/utils/errors"
)

type partidoService struct{}

type partidoServiceInterface interface {
	GetPartidos() (dto.PartidosDto, e.ApiError)
	InsertPartidos(partidoDto dto.PartidoDto) (dto.PartidoDto, e.ApiError)
}

var (
	PartidoService partidoServiceInterface
)

func init() {
	PartidoService = &partidoService{}
}

func (s *partidoService) GetPartidos() (dto.PartidosDto, e.ApiError) {

	var partidos model.Partidos = partidoCliente.GetPartidos()
	var partidosDto dto.PartidosDto
	if len(partidos) == 0 {
		return partidosDto, e.NewBadRequestApiError("partidos no encontrados")
	}
	for _, partido := range partidos {
		var partidoDto dto.PartidoDto
		partidoDto.Id = partido.Id
		partidoDto.IdFase = partido.IdFase
		partidoDto.IdEquipoLocal = partido.IdEquipoLocal
		partidoDto.IdEquipoVisitante = partido.IdEquipoVisitante
		partidoDto.GolesLocal = partido.GolesLocal
		partidoDto.GolesVisitante = partido.GolesVisitante
		partidoDto.Ganador = partido.Ganador

		partidosDto = append(partidosDto, partidoDto)
	}

	return partidosDto, nil
}

InsertPartidos(partidoDto dto.PartidoDto) (dto.PartidoDto, e.ApiError){
	var partido model.Partido

	partido.Id= partidoDto.Id
	partido.IdFase = partidoDto.IdFase
	partido.IdEquipoLocal = partidoDto.IdEquipoLocal
	partido.IdEquipoVisitante = partidoDto.IdEquipoVisitante
	partido.GolesLocal = partidoDto.GolesLocal
	partido.GolesVisitante = partidoDto.GolesVisitante
	partido.Ganador = partidoDto.Ganador

	partido = partidoCliente.InsertPartidos(partido)

	partidoDto.Id = partido.Id
	partidoDto.IdFase = partido.IdFase
	partidoDto.IdEquipoLocal = partido.IdEquipoLocal
	partidoDto.IdEquipoVisitante = partido.IdEquipoVisitante
	partidoDto.GolesLocal = partido.GolesLocal
	partidoDto.GolesVisitante = partido.GolesVisitante
	partidoDto.Ganador = partido.Ganador

	return partidoDto, nil
	
}
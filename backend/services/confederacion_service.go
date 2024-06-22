package services

import (
	confederacionCliente "backend/clients/confederacion"
	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type confederacionService struct{}

type confederacionServiceInterface interface {
	GetConfederaciones() (dto.ConfederacionesDto, e.ApiError)
}

var (
	ConfederacionService confederacionServiceInterface
)

func init() {
	ConfederacionService = &confederacionService{}
}

func (s *confederacionService) GetConfederaciones() (dto.ConfederacionesDto, e.ApiError) {

	var confederaciones model.Confederaciones = confederacionCliente.GetConfederaciones()
	var confederacionesDto dto.ConfederacionesDto
	if len(confederaciones) == 0 {
		return confederacionesDto, e.NewBadRequestApiError("confederacion no encontrada")
	}
	for _, confederacion := range confederaciones {
		var confederacionDto dto.ConfederacionDto
		confederacionDto.Nombre = confederacion.Nombre
		confederacionDto.Logo = confederacion.Logo
		confederacionDto.Id = confederacion.Id

		confederacionesDto = append(confederacionesDto, confederacionDto)
	}

	return confederacionesDto, nil
}

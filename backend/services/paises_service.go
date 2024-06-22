package services

import (
	paisCliente "backend/clients/pais"
	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type paisService struct{}

type paisServiceInterface interface {
	GetPaises() (dto.PaisesDto, e.ApiError)
}

var (
	PaisService paisServiceInterface
)

func init() {
	PaisService = &paisService{}
}

func (s *paisService) GetPaises() (dto.PaisesDto, e.ApiError) {

	var paises model.Paises = paisCliente.GetPaises()
	var paisesDto dto.PaisesDto
	if len(paises) == 0 {
		return paisesDto, e.NewBadRequestApiError("paises no encontrados")
	}
	for _, pais := range paises {
		var paisDto dto.PaisDto
		paisDto.Nombre = pais.Nombre
		paisDto.Id = pais.Id
		paisDto.Bandera = pais.Bandera
		paisDto.IdConfederacion = pais.IdConfederacion

		paisesDto = append(paisesDto, paisDto)
	}

	return paisesDto, nil
}

package services

import (
	faseCliente "back/clients/fase"
	"back/dto"
	"back/model"
	e "back/utils/errors"
)

type faseService struct{}

type faseServiceInterface interface {
	GetFases() (dto.FasesDto, e.ApiError)
}

var (
	FaseService faseServiceInterface
)

func init() {
	FaseService = &faseService{}
}

func (s *faseService) GetFases() (dto.FasesDto, e.ApiError) {

	var fases model.Fases = faseCliente.GetFases()
	var fasesDto dto.FasesDto
	if len(fases) == 0 {
		return fasesDto, e.NewBadRequestApiError("fases no encontradas")
	}
	for _, fase := range fases {
		var faseDto dto.FaseDto
		faseDto.Nombre = fase.Nombre
		faseDto.Id = fase.Id

		fasesDto = append(fasesDto, faseDto)
	}

	return fasesDto, nil
}
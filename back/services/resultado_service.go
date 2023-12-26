package services

import (
	resultadoCliente "back/clients/resultado"
	"back/dto"
	"back/model"
	e "back/utils/errors"
)

type resultadoService struct{}

type resultadoServiceInterface interface {
	GetResultados() (dto.ResultadosDto, e.ApiError)
	InsertResultados(resultadoDto dto.ResultadoDto) (dto.ResultadoDto, e.ApiError)
}

var (
	ResultadoService resultadoServiceInterface
)

func init() {
	ResultadoService = &resultadoService{}
}

func (s *resultadoService) GetResultados() (dto.ResultadosDto, e.ApiError) {

	var resultados model.Resultados = resultadoCliente.GetResultados()
	var resultadosDto dto.ResultadosDto
	if len(resultados) == 0 {
		return resultadosDto, e.NewBadRequestApiError("Resultados de torneos no encontradas")
	}
	for _, resultado := range resultados {
		var resultadoDto dto.ResultadoDto
		resultadoDto.Id = resultado.Id
		resultadoDto.IdEdicionTorneo = resultado.IdEdicionTorneo
		resultadoDto.Campeon = resultado.Campeon
		resultadoDto.Subcampeon = resultado.Subcampeon

		resultadosDto = append(resultadosDto, resultadoDto)
	}

	return resultadosDto, nil
}

func (s *resultadoService) InsertResultados(resultadoDto dto.ResultadoDto) (dto.ResultadoDto, e.ApiError) {

	var resultado model.Resultado

	resultado.Id = resultadoDto.Id
	resultado.IdEdicionTorneo = resultadoDto.IdEdicionTorneo
	resultado.Campeon = resultadoDto.Campeon
	resultado.Subcampeon = resultadoDto.Subcampeon

	resultado = resultadoCliente.InsertResultados(resultado)

	resultadoDto.Id = resultado.Id
	resultadoDto.IdEdicionTorneo = resultado.IdEdicionTorneo
	resultadoDto.Campeon = resultado.Campeon
	resultadoDto.Subcampeon = resultado.Subcampeon

	return resultadoDto, nil

}

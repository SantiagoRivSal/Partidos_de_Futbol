package model

type Fase struct {
	Id              int    `gorm:"primaryKey"`
	Nombre          string `gorm:"type:varchar(350);not null"`
	CantidadEquipos int    `gorm:"type:int(150);not null"`
}
type Fases []Fase

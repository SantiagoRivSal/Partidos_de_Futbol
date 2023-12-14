package model

type Partido struct{
	Id int `gorm:"primaryKey"`
	IdEdicion int `gorm:"type:int(150);not null"`
	IdFase int `gorm:"type:int(150);not null"`
	IdEquipoLocal int `gorm:"type:int(150);not null"`
	IdEquipoVisitante int `gorm:"type:int(150);not null"`
	GolesLocal int `gorm:"type:int(150);not null"`
	GolesVisitante int `gorm:"type:int(150);not null"`
	Ganador int `gorm:"type:int(150);not null"`
}
type Partidos []Partido
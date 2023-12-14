package model

type Partido struct{
	Id int `gorm:"primaryKey"`
	IdEdicionTorneo int `gorm:"type:int(150);not null"`
	IdFase int `gorm:"type:int(150);not null"`
	IdEquipoLocal int `gorm:"type:int(150);not null"`
	IdEquipoVisitante int `gorm:"type:int(150);not null"`
	GolesLocal int `gorm:"type:int(150)"`
	GolesVisitante int `gorm:"type:int(150)"`
	Ganador int `gorm:"type:int(150)"`
}
type Partidos []Partido

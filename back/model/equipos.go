package model

type Equipo struct{
	Id int `gorm:"primaryKey"`
	Nombre string `gorm:"type:varchar(350);not null"`
	Escudo string `gorm:"type:varchar(350);not null"`
	IdPais int `gorm:"type:int(150);not null"`
}
type Equipos []Equipo
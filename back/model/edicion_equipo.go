package model

type EdicionEquipo struct{
	Id int `gorm:"primaryKey"`
	IdEdicionTorneo int `gorm:"type:int(150);not null"`
	IdEquipo int `gorm:"type:int(150);not null"`
}
type EdicionEquipos []EdicionEquipo

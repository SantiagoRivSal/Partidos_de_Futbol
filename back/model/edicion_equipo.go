package model

type EdicionEquipo struct{
	Id int `gorm:"primaryKey"`
	IdEdicion int `gorm:"type:int(150);not null"`
	IdEquipo int `gorm:"type:int(150);not null"`
}
type EdicionEquipos []EdicionEquipo
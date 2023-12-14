package model

type EdicionTorneo struct{
	Id int `gorm:"primaryKey"`
	IdTorneo int `gorm:"type:int(150);not null"`
	Anio int `gorm:"type:int(150);not null"`
	Campeon int `gorm:"type:int(150);not null"`
	Subcampeon int `gorm:"type:int(150);not null"`
}
type EdicionTorneos []EdicionTorneo
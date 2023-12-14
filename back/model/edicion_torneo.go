package model

type EdicionTorneo struct{
	Id int `gorm:"primaryKey"`
	IdTorneo int `gorm:"type:int(150);not null"`
	Anio int `gorm:"type:int(150);not null"`
	Campeon int `gorm:"type:int(150)"`
	Subcampeon int `gorm:"type:int(150)"`
}
type EdicionTorneos []EdicionTorneo
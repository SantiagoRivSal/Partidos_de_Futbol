package model

type EdicionEquipo struct {
	Id              int `gorm:"primaryKey"`
	IdEdicionTorneo int `gorm:"type:int(150);not null"`
	IdEquipo        int `gorm:"type:int(150);not null"`
	//Nombre          string `gorm:"type:varchar(350);not null"`
	//Escudo          string `gorm:"type:varchar(350);not null"`
}
type EdicionEquipos []EdicionEquipo

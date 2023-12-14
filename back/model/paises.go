package model

type Pais struct{
	Id int `gorm:"primaryKey"`
	Nombre string `gorm:"type:varchar(350);not null"`
	Bandera string `gorm:"type:varchar(350);not null"`
	IdConfederacion int `gorm:"type:int(150);not null"`
}
type Paises []Pais
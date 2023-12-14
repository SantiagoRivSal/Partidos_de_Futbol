package model

type Confederacion struct{
	Id int `gorm:"primaryKey"`
	Nombre string  `gorm:"type:varchar(350);not null"`
	Logo string  `gorm:"type:varchar(350);not null"`
}
type Confederaciones []Confederacion
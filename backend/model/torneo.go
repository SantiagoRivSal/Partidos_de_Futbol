package model

type Torneo struct{
	Id int `gorm:"primaryKey"`
	Nombre string `gorm:"type:varchar(350);not null"`
	Logo string `gorm:"type:varchar(350);not null"`
	IdConfederacion int `gorm:"type:int(150);not null"`

}
type Torneos []Torneo
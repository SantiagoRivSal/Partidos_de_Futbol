package model

type Resultado struct {
	Id              int `gorm:"primaryKey"`
	IdEdicionTorneo int `gorm:"type:int(150);not null"`
	Campeon         int `gorm:"type:int(150);not null"`
	Subcampeon      int `gorm:"type:int(150);not null"`
}
type Resultados []Resultado

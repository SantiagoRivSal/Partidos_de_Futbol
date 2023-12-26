package model

type EdicionTorneo struct {
	Id        int    `gorm:"primaryKey"`
	IdTorneo  int    `gorm:"type:int(150);not null"`
	Anio      int    `gorm:"type:int(150);not null"`
	SedeFinal string `gorm:"type:varchar(350);not null"`
}
type EdicionTorneos []EdicionTorneo

package models

type Product struct {
	Name    string  `gorm:"not null"`
	Price   float64 `gorm:"not null"`
	Barcode string  `gorm:"primaryKey"`
	Count   int     `gorm:"not null;default:0"`
}

package models

type Product struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"not null"`
	Price   float64 `gorm:"not null"`
	Barcode string  `gorm:"unique;not null"`
	Count   int     `gorm:"not null"`
}

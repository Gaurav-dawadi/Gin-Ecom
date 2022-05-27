package models

type ProductImage struct {
	Base
	ProductID uint   `json:"product_id" gorm:"not null"`
	Image     string `json:"image" gorm:"not null"`
}

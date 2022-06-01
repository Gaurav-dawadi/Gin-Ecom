package models

type ProductImage struct {
	// Base
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   `json:"product_id" `
	Image     string `json:"image" binding:"required"`
}

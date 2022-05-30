package models

type ProductImage struct {
	// Base
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `form:"product_id" `
	Image     *[]byte `form:"image" binding:"required"`
}

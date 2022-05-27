package models

type Product struct {
	Base
	UserId      uint   `json:"user_id" gorm:"not null"`
	Name        string `json:"name" binding:"required"`
	CategoryID  uint   `json:"category_id" gorm:"not null"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	// Comments    []Comment `json:"comments" gorm:"foreignKey:ProductID;references:ID"`
}

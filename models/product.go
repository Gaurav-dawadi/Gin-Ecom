package models

type Product struct {
	Base
	UserId      uint      `gorm:"not null"`
	Name        string    `json:"name" binding:"required"`
	CategoryID  uint      `gorm:"not null"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	Comments    []Comment `gorm:"foreignKey:ProductID;references:ID"`
}

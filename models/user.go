package models

type User struct {
	Base
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password  string `json:"password" binding:"required" gorm:"not null"`
	// Products  []Product `json:"products" gorm:"foreignKey:UserId;references:ID"`
	// Comments  []Comment `json:"comments" gorm:"foreignKey:UserId;references:ID"`
}

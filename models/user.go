package models

type User struct {
	Base
	Username string    `json:"username"`
	Email    string    `json:"email" binding:"required,email" gorm:"unique"`
	Password string    `form:"password" json:"password" binding:"required"`
	Products []Product `gorm:"foreignKey:UserId;references:ID"`
	Comments []Comment `gorm:"foreignKey:UserId;references:ID"`
}

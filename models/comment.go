package models

type Comment struct {
	Base
	// Parent sql.NullInt64 `json:"parent"`
	UserId    uint   `json:"user_id" gorm:"not null"`
	ProductID uint   `json:"product_id" gorm:"not null"`
	Body      string `json:"body" gorm:"size:1024"`
}

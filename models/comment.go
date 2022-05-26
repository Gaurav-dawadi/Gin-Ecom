package models

type Comment struct {
	Base
	// Parent sql.NullInt64 `json:"parent"`
	UserId    uint   `gorm:"not null"`
	ProductId uint   `gorm:"not null"`
	Body      string `json:"body" gorm:"size:1024"`
}

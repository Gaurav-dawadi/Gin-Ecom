package models

// Todo: Add Parent field which will be self(i.e.) for creating sub-category
type Category struct {
	Base
	// Parent sql.NullInt64 `json:"parent"`
	Title    string    `json:"title"`
	Products []Product `gorm:"foreignKey:CategoryID;references:ID"`
}

package models

type ProductWithImage struct {
	Product Product
	Image   string `json:"image"`
}

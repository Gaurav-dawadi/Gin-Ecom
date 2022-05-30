package models

import (
	"mime/multipart"
)

// *multipart.Form
type ProductWithImage struct {
	Product Product
	File    *multipart.FileHeader `form:"file" json:"file" binding:"required"`
}

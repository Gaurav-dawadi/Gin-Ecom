package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

func GetAllComment() ([]models.Comment, error) {
	var comments []models.Comment
	err := infrastructure.SetupDatabase().Find(&comments).Error
	return comments, err
}

func CreateComment(comment models.Comment) error {
	err := infrastructure.SetupDatabase().Create(comment).Error
	return err
}

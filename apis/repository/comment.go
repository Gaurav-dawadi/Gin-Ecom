package repository

import (
	"go-practice/infrastructure"
	"go-practice/models"
)

type CommentRepository struct {
	database infrastructure.DatabaseSetup
}

func NewCommentRepository(database infrastructure.DatabaseSetup) *CommentRepository {
	return &CommentRepository{
		database: database,
	}
}

func (cr CommentRepository) GetAllComment() ([]models.Comment, error) {
	var comments []models.Comment
	err := cr.database.SetupDatabase().Find(&comments).Error
	return comments, err
}

func (cr CommentRepository) CreateComment(comment models.Comment) error {
	err := cr.database.SetupDatabase().Create(&comment).Error
	return err
}

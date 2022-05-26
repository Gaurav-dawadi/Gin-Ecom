package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

func GetAllComment() ([]models.Comment, error) {
	return repository.GetAllComment()
}

func CreateComment(comment models.Comment) error {
	return repository.CreateComment(comment)
}

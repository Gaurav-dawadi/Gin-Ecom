package services

import (
	"go-practice/apis/repository"
	"go-practice/models"
)

type CommentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (cs CommentService) GetAllComment() ([]models.Comment, error) {
	return cs.commentRepository.GetAllComment()
}

func (cs CommentService) CreateComment(comment models.Comment) error {
	return cs.commentRepository.CreateComment(comment)
}

package services

import (
	"github.com/gmlalfjr/timeline-service/domains/models"
	"github.com/gmlalfjr/timeline-service/repository"
)

type ICommentService interface {
	CreatePostComment(request *models.CommentRequest) (*models.CommentRequest, error)
}

type CommentService struct {
	repository repository.ICommentRepository
}

func NewCommentService(repository repository.ICommentRepository) ICommentService {
	return &CommentService{repository: repository}
}

func (c CommentService) CreatePostComment(request *models.CommentRequest) (*models.CommentRequest, error) {
	//TODO implement me
	return nil, nil
}

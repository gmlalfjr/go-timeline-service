package repository

import (
	"github.com/gmlalfjr/timeline-service/domains/entity"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(timeline *entity.Comment) (*entity.Comment, error)
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{db: db}
}

func (t CommentRepository) CreateComment(timeline *entity.Comment) (*entity.Comment, error) {
	return nil, nil
}

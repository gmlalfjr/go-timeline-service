package services

import (
	"github.com/gmlalfjr/timeline-service/domains/models"
	"github.com/gmlalfjr/timeline-service/repository"
)

type ITimelineService interface {
	CreatePostTimeline(reqeust *models.TimelineRequest) (*models.TimelineRequest, error)
}

type TimelineService struct {
	repository repository.ITimelineRepository
}

func NewTimelineService(repository repository.ITimelineRepository) ITimelineService {
	return &TimelineService{
		repository: repository,
	}
}

func (t TimelineService) CreatePostTimeline(reqeust *models.TimelineRequest) (*models.TimelineRequest, error) {
	return nil, nil
}

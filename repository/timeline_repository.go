package repository

import (
	"github.com/gmlalfjr/timeline-service/domains/entity"
	"gorm.io/gorm"
)

type ITimelineRepository interface {
	CreateTimeline(timeline *entity.Timeline) (*entity.Timeline, error)
}

type TimelineRepository struct {
	db *gorm.DB
}

func (t TimelineRepository) CreateTimeline(timeline *entity.Timeline) (*entity.Timeline, error) {
	return nil, nil
}

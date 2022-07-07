package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gmlalfjr/timeline-service/services"
)

type ITimelineController interface {
	GetTimeline(ctx *gin.Context)
}

type TimelineController struct {
	service services.ITimelineService
}

func NewTimelineController(service services.ITimelineService) ITimelineController {
	return &TimelineController{service: service}
}

func (t TimelineController) GetTimeline(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "success",
	})
	return
}

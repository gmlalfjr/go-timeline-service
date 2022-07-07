package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gmlalfjr/timeline-service/controllers"
)

type Router struct {
	timelineController controllers.ITimelineController
}

func NewRouter(timelineController controllers.ITimelineController) *Router {
	return &Router{
		timelineController: timelineController,
	}
}

func TimelineRouter(router *gin.Engine, controller controllers.ITimelineController) {
	router.GET("/", controller.GetTimeline)
}

func InitRouter(controller controllers.ITimelineController) *gin.Engine {

	router := gin.Default()
	data := NewRouter(controller)
	TimelineRouter(router, data.timelineController)

	return router
}

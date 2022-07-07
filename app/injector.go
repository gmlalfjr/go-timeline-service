package app

import (
	"github.com/gmlalfjr/timeline-service/controllers"
	"github.com/gmlalfjr/timeline-service/repository"
	"github.com/gmlalfjr/timeline-service/services"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func TimelineInjector(conn *gorm.DB, server *grpc.Server) controllers.ITimelineController {
	newRepo := repository.NewTimelineRepository(conn)
	newService := services.NewTimelineService(newRepo)
	newController := controllers.NewTimelineController(newService)
	RegisterGrpcRoutes(newService, server)
	
	return newController
}

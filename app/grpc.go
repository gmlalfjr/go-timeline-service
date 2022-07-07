package app

import (
	"github.com/gmlalfjr/timeline-service/controllers"
	pb "github.com/gmlalfjr/timeline-service/gen/proto/timeline"
	"github.com/gmlalfjr/timeline-service/services"
	"google.golang.org/grpc"
)

type AllServerGrpc struct {
	timelineService services.ITimelineService
}

func (serve AllServerGrpc) AllGrpc() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterTimelineServer(server, controllers.NewTimelineServerGrpc(pb.UnimplementedTimelineServer{}, serve.timelineService))
	return server
}

func RegisterGrpcRoutes(service services.ITimelineService, server *grpc.Server) {
	pb.RegisterTimelineServer(server, controllers.NewTimelineServerGrpc(pb.UnimplementedTimelineServer{}, service))
}

package controllers

import (
	"context"
	"fmt"
	"github.com/gmlalfjr/timeline-service/domains/models"
	pb "github.com/gmlalfjr/timeline-service/gen/proto/timeline"
	"github.com/gmlalfjr/timeline-service/services"
)

type TimelineServer struct {
	pb.UnimplementedTimelineServer
	service services.ITimelineService
}

func NewTimelineServerGrpc(unimplementedTimelineServer pb.UnimplementedTimelineServer, service services.ITimelineService) pb.TimelineServer {
	return &TimelineServer{UnimplementedTimelineServer: unimplementedTimelineServer, service: service}
}

func (s TimelineServer) CreateTimeline(ctx context.Context, n *pb.TimelineRequest) (*pb.TimelineResponse, error) {
	fmt.Println("kena	 mari")
	_, err := s.service.CreatePostTimeline(&models.TimelineRequest{})
	if err != nil {
		return nil, err
	}
	return &pb.TimelineResponse{
		Status:  "woi",
		Message: "woi",
		Data: &pb.TimelineResponseData{
			PostText:  "wkkw",
			IsPrivate: "wkwk2",
		},
	}, nil

}

func (s TimelineServer) GetTimelinebyId(ctx context.Context, n *pb.GetTimelineRequest) (*pb.GetTimelineResponse, error) {
	return &pb.GetTimelineResponse{Id: n.Id}, nil
}

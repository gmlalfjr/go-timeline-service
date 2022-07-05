package main

import (
	"context"
	"fmt"
	pb "github.com/gmlalfjr/timeline-service/grpc-gateway/gen/proto/timeline"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTimelineClient(conn)
	resp, err := client.CreateTimeline(context.Background(), &pb.TimelineRequest{
		PostText:  "woii",
		IsPrivate: "woii2",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.PostText)
}

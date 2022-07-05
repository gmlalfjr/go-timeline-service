package main

import (
	"context"
	"fmt"
	pb "github.com/gmlalfjr/timeline-service/gen/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type Server struct {
	pb.UnimplementedTimelineServer
}

func (s Server) CreateTimeline(ctx context.Context, n *pb.TimelineRequest) (*pb.TimelineResponse, error) {
	fmt.Println("kena mari bossku")
	return &pb.TimelineResponse{
		PostText:  n.PostText,
		IsPrivate: n.IsPrivate,
	}, nil
}

func main() {
	go func() {
		mux := runtime.NewServeMux()
		pb.RegisterTimelineHandlerServer(context.Background(), mux, &Server{})
		err := http.ListenAndServe("localhost:8081", mux)
		if err != nil {
			fmt.Println(err, "log error")
		}
		fmt.Println("success")
		//log.Fatal()
	}()
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTimelineServer(grpcServer, &Server{})

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
	log.Printf("running on port %v", 8080)
}

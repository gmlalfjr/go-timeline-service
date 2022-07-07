package app

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/gmlalfjr/timeline-service/gen/proto/timeline"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"time"
)

func Connection() *gorm.DB {
	dsn := "host=localhost user=postgres password=12341234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
	sqlDB, err := client.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	return client
}

func RunServer() *gin.Engine {
	mux := runtime.NewServeMux()
	err := pb.RegisterTimelineHandlerFromEndpoint(
		context.Background(), mux,
		"localhost:8081",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	r.Use(gin.Logger())

	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}
	route, grpcServer := Serve()
	go func() {
		err = grpcServer.Serve(listen)
		if err != nil {
			log.Println(err)
		}
	}()
	route.Group("grpc/*{grpc_gateway}").Any("", gin.WrapH(mux))

	return route
}

func Serve() (*gin.Engine, *grpc.Server) {
	conn := Connection()
	server := grpc.NewServer()
	timeline := TimelineInjector(conn, server)
	route := InitRouter(timeline)

	return route, server
}

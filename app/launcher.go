package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gmlalfjr/timeline-service/controllers"
	pb "github.com/gmlalfjr/timeline-service/gen/proto/timeline"
	"github.com/gmlalfjr/timeline-service/repository"
	"github.com/gmlalfjr/timeline-service/services"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func Connection() *gorm.DB {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	conf.Get("DB_USERNAME"),
	//	conf.Get("DB_PASSWORD"),
	//	conf.Get("PORT"),
	//	conf.Get("DATABASE"),
	//)
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

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in allowed header, don't send the header
	return s, false
}

type Server struct {
	pb.UnimplementedTimelineServer
}

func (s Server) CreateTimeline(ctx context.Context, n *pb.TimelineRequest) (*pb.TimelineResponse, error) {
	return &pb.TimelineResponse{
		PostText:  n.PostText,
		IsPrivate: n.IsPrivate,
	}, nil
}

func RunServer() *gin.Engine {
	mux := runtime.NewServeMux(
		// convert header in response(going from gateway) from metadata received.
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			md := metadata.Pairs("auth", header)
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			//creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		}))

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
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterTimelineServer(grpcServer, &Server{})

		err = grpcServer.Serve(listen)
		if err != nil {
			log.Println(err)
		}
	}()

	conn := Connection()
	newRepo := repository.NewTimelineRepository(conn)
	newService := services.NewTimelineService(newRepo)
	newController := controllers.NewTimelineController(newService)
	r.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(mux))
	r.GET("/timeline", newController.GetTimeline)

	return r
}

package grpc_port

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"users-service-cqrs/internal/app"
)

type GrpcServer struct {
	s *grpc.Server
}

func NewServer() *GrpcServer {
	return &GrpcServer{}
}

func (gs *GrpcServer) Start(port int, application *app.App) {
	gs.s = grpc.NewServer(grpc.EmptyServerOption{})
	RegisterUsersServiceServer(gs.s, Handlers{application: application})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("unable to start the grpc server: %v", err)
	}

	go func() {
		log.Printf("The grpc server is running at port %d\n", port)
		log.Println("The server stopped due to: ", gs.s.Serve(listener))
	}()
}

func (gs *GrpcServer) Stop() {
	gs.s.GracefulStop()
	log.Println("The grpc server has stopped gracefully")
}

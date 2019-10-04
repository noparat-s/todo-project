package grpc

import (
	"context"
	"log"
	"net"
	"todo/pkg/server"

	"google.golang.org/grpc"
)

type grpcServer struct {
	server *grpc.Server
	li     net.Listener
}

// NewGRPCServer create gRPC server
func NewGRPCServer() server.EntryPoint {
	li, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic("cannot binding grpc server")
	}
	return &grpcServer{
		server: grpc.NewServer(),
		li:     li,
	}
}

func (s *grpcServer) Start(ctx context.Context) error {
	log.Print("start gRPC server")
	return s.server.Serve(s.li)
}

func (s *grpcServer) ShutDown(ctx context.Context) {
	defer log.Println("stopping grpc server")
	go func() {
		s.server.GracefulStop()
		<-ctx.Done()
	}()
}

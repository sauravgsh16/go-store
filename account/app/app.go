package app

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sauravgsh16/go-store/account/controller"
	"github.com/sauravgsh16/go-store/account/pb"
	"github.com/sauravgsh16/go-store/account/service"
)

// RunGRPC server
func RunGRPC(ctx context.Context, s service.Service, port string) error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAccountServiceServer(grpcServer, &controller.AccGrpc{
		Serv: s,
	})

	// Reflections for register server
	reflection.Register(grpcServer)

	// graceful shutdown
	c := make(chan os.Signal, 1)

	// Will relay all signal from os.Interrupt to chan c
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("shutting down gRPC server....")
			grpcServer.GracefulStop()

			// Send done signal to context
			<-ctx.Done()
		}

	}()

	return grpcServer.Serve(ln)
}

package main

import (
	"fmt"
	"log"
	"net"

	grpcdelivery "github.com/syafdia/sb/chal002/internal/delivery/grpc"
	"github.com/syafdia/sb/chal002/internal/di"
	"google.golang.org/grpc"
)

func main() {
	// Setup module.
	appModule := di.NewAppModule()
	repoModule := di.NewRepoModule(appModule)
	useCaseModule := di.NewUseCaseModule(repoModule)

	// Setup GRPC handler.
	movieHandler := grpcdelivery.NewMovieHandler(
		useCaseModule.FindOneMovieUseCase,
		useCaseModule.FindMultipleMoviesUseCase,
	)

	// Setup GRPC Server.
	server := grpc.NewServer()
	grpcdelivery.RegisterMovieServiceServer(server, movieHandler)

	address := "0.0.0.0:8002"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("GRPC server is listening on %v ...", address)

	server.Serve(listener)
}

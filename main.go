package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/ssshekhu53/user-detail-management/grpc"
	handlerUser "github.com/ssshekhu53/user-detail-management/handler/user"
	serviceUser "github.com/ssshekhu53/user-detail-management/service/user"
	storeUser "github.com/ssshekhu53/user-detail-management/store/user"
)

func main() {
	grpcPortEnv := os.Getenv("GRPC_PORT")

	_, err := strconv.Atoi(grpcPortEnv)
	if err != nil {
		grpcPortEnv = "9000"
	}

	grpcPort := fmt.Sprintf(":%s", grpcPortEnv)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userStore := storeUser.New()
	userSvc := serviceUser.New(userStore)
	userHandler := handlerUser.New(userSvc)

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, userHandler)

	log.Printf("Starting gRPC server on port %s", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

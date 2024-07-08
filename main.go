package main

import (
	"fmt"
	"github.com/ssshekhu53/user-detail-management/interceptor"
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
	logger := log.New(os.Stdout, "user-detail-management ", log.Ldate|log.Ltime)

	grpcPortEnv := os.Getenv("GRPC_PORT")

	_, err := strconv.Atoi(grpcPortEnv)
	if err != nil {
		grpcPortEnv = "9000"
	}

	grpcPort := fmt.Sprintf(":%s", grpcPortEnv)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}

	userStore := storeUser.New()
	userSvc := serviceUser.New(userStore)
	userHandler := handlerUser.New(userSvc)

	loggingInterceptor := interceptor.NewLoggingInterceptor(logger)

	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor.UnaryLoggingInterceptor))

	pb.RegisterUserServiceServer(s, userHandler)

	logger.Printf("Starting gRPC server on port %s", grpcPort)
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}

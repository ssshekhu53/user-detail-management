package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type loggingInterceptor struct {
	logger *log.Logger
}

func NewLoggingInterceptor(logger *log.Logger) *loggingInterceptor {
	return &loggingInterceptor{logger: logger}
}

func (l *loggingInterceptor) UnaryLoggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	h, err := handler(ctx, req)

	end := time.Now()

	l.logger.Printf("RPC: %s, Duration: %s, Error: %v", info.FullMethod, end.Sub(start), err)

	return h, err
}

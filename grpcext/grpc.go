package grpcext

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewInsecureClient(cfg *Config) (*grpc.ClientConn, error) {
	return grpc.Dial(cfg.ClientAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}

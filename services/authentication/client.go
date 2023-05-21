package authentication

import (
	"github.com/ungame/go-signup/grpcext"
	"github.com/ungame/go-signup/pb/auth"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	conn   *grpc.ClientConn
	client auth.AuthenticationServiceClient
}

func NewServiceClient(cfg *grpcext.Config) (ServiceClient, error) {
	conn, err := grpcext.NewInsecureClient(cfg)
	if err != nil {
		return ServiceClient{}, err
	}
	return ServiceClient{
		client: auth.NewAuthenticationServiceClient(conn),
	}, nil
}

func (c ServiceClient) Client() auth.AuthenticationServiceClient {
	return c.client
}

func (c ServiceClient) Close() error {
	return c.conn.Close()
}

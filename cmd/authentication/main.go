package main

import (
	"context"
	"flag"
	"github.com/ungame/go-signup/db/mysqlext"
	"github.com/ungame/go-signup/httpext"
	"github.com/ungame/go-signup/logext"
	"github.com/ungame/go-signup/pb/auth"
	"github.com/ungame/go-signup/services/authentication"
	"github.com/ungame/go-signup/services/authentication/repository"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func init() {
	authentication.LoadJSONConfigFromFlags(flag.CommandLine)
	mysqlext.LoadJSONConfigFromFlags(flag.CommandLine)
	flag.Parse()
}

func main() {
	logger, err := logext.New("authentication")
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Close()

	var (
		mysqlConfig = mysqlext.GetDefaultMySQLConfig()
		authConfig  = authentication.GetConfigs()
		healthCheck = httpext.NewHealthChecker(8180)
	)

	listener, err := net.Listen("tcp", authConfig.ServerAddr())
	if err != nil {
		logger.Fatal("unable to create listener for grpc: %s", err.Error())
	}

	mysqlCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	mysqlConn := mysqlext.New(mysqlCtx, mysqlConfig)
	defer func() { _ = mysqlConn.Close() }()

	authenticationUsersRepository := repository.NewAuthenticationUsersRepository(mysqlConn)
	authenticationUsersService := authentication.NewAuthenticationService(authenticationUsersRepository, logger)

	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	auth.RegisterAuthenticationServiceServer(grpcServer, authenticationUsersService)

	go func() {
		healthCheck.Start()
	}()

	log.Printf("Authentication GRPC service listening on [::]:%d\n\n", authConfig.Port)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("grpc serve error: %s", err.Error())
	}
}

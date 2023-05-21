package authentication

import (
	"context"
	"github.com/google/uuid"
	"github.com/ungame/go-signup/logext"
	"github.com/ungame/go-signup/pb/auth"
	"github.com/ungame/go-signup/security"
	"github.com/ungame/go-signup/services/authentication/entities"
	"github.com/ungame/go-signup/services/authentication/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type authenticationService struct {
	authenticationUsersRepository repository.AuthenticationUsersRepository
	logger                        logext.Logger
}

func NewAuthenticationService(authenticationUsersRepository repository.AuthenticationUsersRepository, logger logext.Logger) auth.AuthenticationServiceServer {
	return &authenticationService{
		authenticationUsersRepository: authenticationUsersRepository,
		logger:                        logger,
	}
}

func (s authenticationService) CreateAuthentication(ctx context.Context, request *auth.CreateAuthenticationRequest) (*auth.AuthenticationUser, error) {

	password, err := security.GeneratePassword(request.Password)
	if err != nil {
		s.logger.Error("CreateAuthentication: failed to generate password. Error=%s", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	entity := &entities.DbAuthenticationUser{
		Id:        uuid.NewString(),
		Email:     request.Email,
		Username:  request.Username,
		Password:  password,
		Phone:     request.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.authenticationUsersRepository.Create(ctx, entity)
	if err != nil {
		s.logger.Error("CreateAuthentication: failed to save entity. Error=%s", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &auth.AuthenticationUser{
		Id:        entity.Id,
		Email:     entity.Email,
		Username:  entity.Username,
		Phone:     entity.Phone,
		CreatedAt: timestamppb.New(entity.CreatedAt),
		UpdatedAt: timestamppb.New(entity.UpdatedAt),
	}, nil
}

func (s authenticationService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {

	existing, err := s.authenticationUsersRepository.Get(ctx, "username", request.Username)
	if err != nil {
		return nil, err
	}

	err = security.VerifyPassword(request.Password, existing.Password)
	if err != nil {
		s.logger.Error("Login: failed to verify password. Error=%s", err.Error())
		return nil, err
	}

	token, err := security.NewToken(existing.Id)
	if err != nil {
		s.logger.Error("Login: failed to create token. Error=%s", err.Error())
		return nil, err
	}

	return &auth.LoginResponse{
		Token: token,
		AuthenticationUser: &auth.AuthenticationUser{
			Id:        existing.Id,
			Email:     existing.Email,
			Username:  existing.Username,
			Phone:     existing.Phone,
			CreatedAt: timestamppb.New(existing.CreatedAt),
			UpdatedAt: timestamppb.New(existing.UpdatedAt),
		},
	}, nil
}

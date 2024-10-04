package authServer

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	SigninDto "github.com/MoreWiktor/auth.sso.back/internal/grpc/auth/dtos"
	authv1 "github.com/MoreWiktor/go.sso.proto/auth"
)

type serverAPI struct {
	authv1.AuthServer
	auth Auth
}

type Auth interface {
	Signin(
		ctx context.Context,
		dto *SigninDto.Result,
	) (token string, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	authv1.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Signin(
	ctx context.Context,
	in *authv1.SigninRequest,
) (*authv1.SigninResponse, error) {

	dto, err := SigninDto.RequestValidator(in)

	if err != nil {
		return nil, err
	}

	token, err := s.auth.Signin(ctx, dto)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &authv1.SigninResponse{Token: token}, nil
}

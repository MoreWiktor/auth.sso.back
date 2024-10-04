package auth

import (
	"context"
	"errors"
	"log/slog"
	"time"

	models "github.com/MoreWiktor/auth.sso.back/internal/domain"
	SigninRequestDto "github.com/MoreWiktor/auth.sso.back/internal/grpc/auth/dtos"
)

type Auth struct {
	log         *slog.Logger
	tokenTTL    time.Duration
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserToken interface {
    SaveToken(
        ctx context.Context,
        payload models.Token,
    ) (uid int64, err error)
}

type TokenProvider interface {  
    User(ctx context.Context, login string) (models.Token, error)  
}

type TokenStorage interface {
	TokenByLogin(ctx context.Context, login string) (models.Token, error)
    CreateToken(ctx context.Context, payload models.Token) (uid int64, err error)
}

func New(
	log *slog.Logger,
	tokenTTL time.Duration,
) *Auth {
	return &Auth {
		log,
		tokenTTL,
	}
}

// func (a *Auth) Signin () {}

func (a *Auth) Signin(context.Context, *SigninRequestDto.Result) (string, error) {

	return "token", nil
}

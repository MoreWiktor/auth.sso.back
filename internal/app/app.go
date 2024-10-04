package app

import (
	"log/slog"
	"time"

	grpcApp "github.com/MoreWiktor/auth.sso.back/internal/app/grpc"
	"github.com/MoreWiktor/auth.sso.back/internal/services/auth"
)

type App struct {
	GRPCServer *grpcApp.App
}

func New(
	log *slog.Logger,
	port int,
	storagePath string,
	tokenTTL time.Duration,
) *App {

	authService := auth.New(log, tokenTTL)

	grpcApp := grpcApp.New(log, authService, port)

	return &App{
		GRPCServer: grpcApp,
	}
}

func (a *App) Signin() {
	panic("impl me")
}

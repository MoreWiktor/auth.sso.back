package grpcApp

import (
	"fmt"
	"log/slog"
	"net"

	authServer "github.com/MoreWiktor/auth.sso.back/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	log *slog.Logger
	gRPCServer *grpc.Server
	port int
}

func New(
	log *slog.Logger,
	authService authServer.Auth,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	authServer.Register(gRPCServer, authService)

	return &App{
		log,
		gRPCServer,
		port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcApp.Run"
	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	listner, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info(fmt.Sprintf("gRPC server listen %d port...", a.port))

	if err := a.gRPCServer.Serve(listner); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcApp.Stop"

	a.log.With(slog.String("op", op)).Info("stoping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}

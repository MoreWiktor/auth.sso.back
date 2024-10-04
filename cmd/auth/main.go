package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/MoreWiktor/auth.sso.back/internal/app"
	"github.com/MoreWiktor/auth.sso.back/internal/config"
	"github.com/MoreWiktor/auth.sso.back/internal/logger"
)

func main() {  
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL);

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()

	log.Info("Application stopped")

}

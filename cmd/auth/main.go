package main

import (
	"github.com/MoreWiktor/auth.sso.back/internal/config"
	"github.com/MoreWiktor/auth.sso.back/internal/logger"
)

func main() {  
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
}

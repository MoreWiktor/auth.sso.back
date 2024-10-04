# Makefile
run:
	CONFIG_PATH="./config/config_local.yaml" go run ./cmd/auth/main.go

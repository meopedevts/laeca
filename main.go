package main

import (
	"github.com/meopedevts/laeca/cmd"
	"github.com/meopedevts/laeca/config"
	"github.com/meopedevts/laeca/internal/logger"
)

func main() {
	logger.InitLogger()
	cfg := config.LoadConfig()

	cmd.StartServer(cfg)
}

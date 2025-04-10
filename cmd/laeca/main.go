package main

import (
	"flag"
	"log/slog"

	"github.com/meopedevts/laeca/config"
	"github.com/meopedevts/laeca/internal/logger"
	"github.com/meopedevts/laeca/internal/server"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	logger.InitLogger(slog.LevelInfo, *debug)

	args := flag.Args()
	if len(args) == 0 {
		logger.Fatal("error on starting server", "reason", "no command provided")
	}

	cfg := config.LoadConfig()
	cmd := args[0]

	switch cmd {
	case "start":
		server := server.New(cfg)
		server.Start()

	default:
		logger.Fatal("Unknown command", "command", cmd)
	}
}

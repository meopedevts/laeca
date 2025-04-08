package cmd

import (
	"net"
	"strconv"

	"github.com/meopedevts/laeca/config"
	"github.com/meopedevts/laeca/internal/logger"
	"github.com/meopedevts/laeca/proxy"
)

func StartServer(cfg *config.Config) {
	l, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.Listen))
	if err != nil {
		logger.Fatal("failed to start TCP listener: %v", err)
	}
	defer l.Close()

	logger.Warn("LaECa started on %s", l.Addr().String())
	logger.Warn("Available backends: %s", cfg.Upstream)
	logger.Warn("Protocol: %s", cfg.Protocol)
	logger.Warn("Algorithm: %s", cfg.Algorithm)

	p := proxy.New()

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Error("failed to accept connection: %v", err)
			continue
		}

		go p.Handle(conn)
	}
}

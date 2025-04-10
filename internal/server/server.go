package server

import (
	"net"
	"strconv"

	"github.com/meopedevts/laeca/config"
	"github.com/meopedevts/laeca/internal/logger"
	"github.com/meopedevts/laeca/proxy"
)

type Server struct {
	config *config.Config
}

func New(config *config.Config) *Server {
	server := &Server{
		config: config,
	}

	return server
}

func (s *Server) Start() {
	log := logger.With("component", "server")

	addr := ":" + strconv.Itoa(s.config.Listen)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal("failed to start TCP listener: %v", err)
	}
	defer listener.Close()

	// logger.Warn("LaECa started on %s", l.Addr().String())
	// logger.Warn("Available backends: %s", s.config.Upstream)
	// logger.Warn("Protocol: %s", s.config.Protocol)
	// logger.Warn("Algorithm: %s", s.config.Algorithm)
	log.Info("LaECa started",
		"address", listener.Addr().String(),
		"protocol", s.config.Protocol,
		"algorithm", s.config.Algorithm,
		"upstreams", s.config.Upstream,
	)

	p := proxy.New()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("failed to accept connection", "error", err)
			continue
		}

		go p.Handle(conn)
	}
}

package proxy

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/meopedevts/laeca/balancer"
	"github.com/meopedevts/laeca/config"
	"github.com/meopedevts/laeca/internal/logger"
)

type Proxy struct {
	cfg        *config.Config
	balancer   balancer.Balancer
	httpClient *http.Client
}

func New() *Proxy {
	cfg := config.LoadConfig()
	return &Proxy{
		cfg:        cfg,
		balancer:   balancer.New(cfg),
		httpClient: &http.Client{},
	}
}

func (p *Proxy) Handle(conn net.Conn) {
	defer conn.Close()

	remote := conn.RemoteAddr().String()
	log := logger.With("component", "proxy", "remote", remote)

	backend, err := p.balancer.Next()
	if err != nil {
		log.Error("failed to get next backend", "error", err)
		return
	}

	switch p.cfg.Protocol {
	case "tcp":
		log.Info("handling TCP request")
		if err := p.forwardTCP(conn, backend); err != nil {
			log.Error("failed to forward TCP connection", "error", err)
		}

	case "http":
		log.Info("handling HTTP request")
		if err := p.forwardHTTP(conn, backend); err != nil {
			log.Error("failed to forward HTTP connection", "error", err)
		}

	default:
		log.Warn("unknown protocol, falling back to TCP")
		if err := p.forwardTCP(conn, backend); err != nil {
			log.Error("failed to forward TCP connection", "error", err)
		}
	}
}

func (p *Proxy) forwardTCP(downstream net.Conn, backend string) error {
	upstream, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed to connect to backend %s: %v", backend, err)
	}
	defer upstream.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	buf := make([]byte, 32*1024)

	go func() {
		defer wg.Done()
		io.CopyBuffer(upstream, downstream, buf)
	}()

	go func() {
		defer wg.Done()
		io.CopyBuffer(downstream, upstream, buf)
	}()

	wg.Wait()
	return nil
}

func (p *Proxy) forwardHTTP(downstream net.Conn, backend string) error {
	req, err := http.ReadRequest(bufio.NewReader(downstream))
	if err != nil {
		return fmt.Errorf("failed to read HTTP request: %w", err)
	}
	req.RequestURI = ""
	req.URL.Scheme = "http"
	req.URL.Host = backend

	res, err := p.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to forward HTTP request: %w", err)
	}
	defer res.Body.Close()

	if err := res.Write(downstream); err != nil {
		return fmt.Errorf("failed to write HTTP response: %w", err)
	}

	return nil
}

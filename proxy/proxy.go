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
	cfg      *config.Config
	balancer balancer.Balancer

	httpClient *http.Client
}

func New() *Proxy {
	cfg := config.LoadConfig()
	balancer := balancer.New(cfg)

	return &Proxy{
		cfg:      cfg,
		balancer: balancer,

		httpClient: &http.Client{},
	}
}

func (p *Proxy) Handle(conn net.Conn) {
	backend, err := p.balancer.Next()
	if err != nil {
		logger.Error("failed to get next backend %v", err)
	}

	switch p.cfg.Protocol {
	case "tcp":
		if err := p.forwardTCP(&conn, backend); err != nil {
			logger.Error("failed to forward TCP connection: %v", err)
		}

	case "http":
		if err := p.forwardHTTP(&conn, backend); err != nil {
			logger.Error("failed to forward HTTP connection: %v", err)
		}

	default:
		if err := p.forwardTCP(&conn, backend); err != nil {
			logger.Error("failed to forward TCP connection: %v", err)
		}
	}
}

func (p *Proxy) forwardTCP(downstream *net.Conn, backend string) error {
	upstream, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed to connect to backend %s: %v", backend, err)
	}
	defer upstream.Close()
	defer (*downstream).Close()

	logger.Info("New TCP request from %s to %s", (*downstream).RemoteAddr().String(), backend)

	var wg sync.WaitGroup
	wg.Add(2)

	buf := make([]byte, 32*1024)

	go func() {
		defer wg.Done()
		io.CopyBuffer(upstream, *downstream, buf)
	}()

	go func() {
		defer wg.Done()
		io.CopyBuffer(*downstream, upstream, buf)
	}()

	wg.Wait()
	return nil
}

func (p *Proxy) forwardHTTP(downstream *net.Conn, backend string) error {
	defer (*downstream).Close()

	req, err := http.ReadRequest(bufio.NewReader(*downstream))
	if err != nil {
		return fmt.Errorf("failed to read HTTP request: %v", err)
	}
	req.RequestURI = ""
	req.URL.Scheme = "http"
	req.URL.Host = backend

	logger.Info("New HTTP request from %s to %s", (*downstream).RemoteAddr().String(), backend)

	res, err := p.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to forward HTTP request: %v", err)
	}
	defer res.Body.Close()

	if err := res.Write(*downstream); err != nil {
		return fmt.Errorf("failed to write HTTP response: %v", err)
	}

	return nil
}

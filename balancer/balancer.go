package balancer

import (
	"github.com/meopedevts/laeca/config"
)

type Balancer interface {
	Next() (string, error)
}

func New(cfg *config.Config) Balancer {
	switch cfg.Algorithm {
	case "round-robinrobin":
		return newRoundRobin(cfg.Upstream)
	default:
		return newRoundRobin(cfg.Upstream)
	}
}

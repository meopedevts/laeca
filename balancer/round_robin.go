package balancer

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/meopedevts/laeca/config"
	"github.com/meopedevts/laeca/internal/logger"
)

type roundRobinBalancer struct {
	Balancer

	upstream []config.Upstream

	counter atomic.Int32
}

func newRoundRobin(upstream []config.Upstream) *roundRobinBalancer {
	return &roundRobinBalancer{
		upstream: upstream,
	}
}

func (b *roundRobinBalancer) Next() (string, error) {
	logger.Debug("%s", "Counter: "+strconv.Itoa(int(b.counter.Load())))
	server := b.upstream[b.counter.Load()%int32(len(b.upstream))]
	b.counter.Add(1)
	return fmt.Sprintf("%s:%s", server.Url, server.Port), nil
}

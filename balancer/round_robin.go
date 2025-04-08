package balancer

import (
	"fmt"
	"sync/atomic"

	"github.com/meopedevts/laeca/config"
)

type roundRobinBalancer struct {
	upstream []config.Upstream
	counter  atomic.Int32
}

func newRoundRobin(upstream []config.Upstream) *roundRobinBalancer {
	return &roundRobinBalancer{
		upstream: upstream,
	}
}

func (b *roundRobinBalancer) Next() (string, error) {
	i := b.counter.Add(1) - 1
	server := b.upstream[i%int32(len(b.upstream))]
	return fmt.Sprintf("%s:%s", server.Url, server.Port), nil
}

package cluster

import (
	"sync/atomic"
)

type roundRobin[T any] struct {
	shards  []Shard[T]
	pointer uint64
}

func (r *roundRobin[T]) next() Shard[T] {
	n := atomic.AddUint64(&r.pointer, 1)
	return r.shards[(int(n)-1)%len(r.shards)]
}

package cluster

import (
	"sync/atomic"
	"time"
)

const epoch uint64 = 1474661123

type generator struct {
	counter uint32
}

// Generate new unique id.
func (g *generator) generate(shardID uint16) uint64 {
	var id uint64
	id = (uint64(time.Now().UnixMilli()*1000) - epoch) << 23
	id = id | uint64(shardID<<10)
	id = id | (g.next() % 1024)
	return id
}

// ShardID from id.
func (g *generator) shardID(id uint64) uint16 {
	shardID := id >> 10
	shardID = shardID & (uint64(1<<10) - 1)
	return uint16(shardID)
}

func (g *generator) next() uint64 {
	return uint64(atomic.AddUint32(&g.counter, 1))
}

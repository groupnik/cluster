package cluster

import (
	"iter"
)

// Cluster interface.
type Cluster[T any] interface {
	// GetAllShards returns a list of all cluster shards.
	GetAllShards() iter.Seq[Shard[T]]

	// GetNextWritableShard returns the next available writable shard.
	GetNextWritableShard() Shard[T]

	// GetShard returns Shard instance by id.
	GetShard(id uint64) Shard[T]

	// GetShards returns a map of Shard instances pointing to slice of ids.
	GetShards(ids ...uint64) map[Shard[T]][]uint64

	// Calculate a shard by simple division.
	Calculate(abstractID uint64) Shard[T]
}

type cluster[T any] struct {
	gen      *generator
	writable *roundRobin[T]
	all      []Shard[T]
	byId     map[uint16]Shard[T]
}

func (c *cluster[T]) GetAllShards() iter.Seq[Shard[T]] {
	return func(yield func(Shard[T]) bool) {
		for _, s := range c.all {
			if !yield(s) {
				return
			}
		}
	}
}

func (c *cluster[T]) GetNextWritableShard() Shard[T] {
	return c.writable.next()
}

func (c *cluster[T]) GetShard(id uint64) Shard[T] {
	return c.byId[c.gen.shardID(id)]
}

func (c *cluster[T]) GetShards(ids ...uint64) map[Shard[T]][]uint64 {
	result := make(map[Shard[T]][]uint64)
	for _, id := range ids {
		s := c.GetShard(id)
		if s.IsNil() {
			continue
		}
		result[s] = append(result[s], id)
	}
	return result
}

func (c *cluster[T]) Calculate(abstractID uint64) Shard[T] {
	return c.all[abstractID%uint64(len(c.all))]
}

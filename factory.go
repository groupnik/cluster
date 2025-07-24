package cluster

import (
	"slices"
)

func New[T any](configs ...ShardConfig[T]) Cluster[T] {
	if len(configs) == 0 {
		panic("cannot initialize cluster with empty configs")
	}
	slices.SortFunc[[]ShardConfig[T], ShardConfig[T]](configs, func(a, b ShardConfig[T]) int {
		if a.ID > b.ID {
			return 1
		}
		return -1
	})
	c := &cluster[T]{
		all: make([]Shard[T], len(configs)),
		writable: &roundRobin[T]{
			shards:  make([]Shard[T], 0, len(configs)),
			pointer: 0,
		},
		byId: make(map[uint16]Shard[T]),
		gen:  new(generator),
	}

	for i, config := range configs {
		if _, ok := c.byId[config.ID]; ok {
			panic("duplicate shard id")
		}
		s := newShard[T](config.ID, config.DB, c.gen)
		c.all[i] = s
		c.byId[config.ID] = s
		if !config.ReadOnly {
			c.writable.shards = append(c.writable.shards, s)
		}
	}
	return c
}

package cluster

// Shard interface.
type Shard[T any] interface {
	// ID returns the shard id.
	ID() uint16

	// T returns T assigned to shard.
	T() T

	// Next returns generated id.
	Next() uint64

	// IsNil checks if the shard is nil.
	IsNil() bool
}

func newShard[T any](id uint16, db T, gen *generator) Shard[T] {
	return &shard[T]{id, db, gen}
}

type shard[T any] struct {
	id  uint16
	t   T
	gen *generator
}

func (s *shard[T]) ID() uint16 {
	return s.id
}

func (s *shard[T]) T() T {
	return s.t
}

func (s *shard[T]) Next() uint64 {
	return s.gen.generate(s.id)
}

func (s *shard[T]) IsNil() bool {
	return s == nil
}

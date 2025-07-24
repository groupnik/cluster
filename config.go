package cluster

type ShardConfig[T any] struct {
	ID       uint16
	DB       T
	ReadOnly bool
}

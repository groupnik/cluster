package cluster

import (
	"reflect"
	"testing"
)

func Test_newShard(t *testing.T) {
	gen := new(generator)
	type args[T any] struct {
		id  uint16
		db  T
		gen *generator
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want Shard[T]
	}
	tests := []testCase[int]{
		{"check new shard 1", args[int]{1, 23, gen}, &shard[int]{1, 23, gen}},
		{"check new shard 2", args[int]{2, 24, gen}, &shard[int]{2, 24, gen}},
		{"check new shard 3", args[int]{3, 25, gen}, &shard[int]{3, 25, gen}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newShard(tt.args.id, tt.args.db, tt.args.gen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shard_ID(t *testing.T) {
	gen := new(generator)
	type testCase[T any] struct {
		name string
		s    shard[T]
		want uint16
	}
	tests := []testCase[int]{
		{"check id 1", shard[int]{1, 23, gen}, 1},
		{"check id 2", shard[int]{2, 23, gen}, 2},
		{"check id 3", shard[int]{3, 23, gen}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shard_IsNil(t *testing.T) {
	gen := new(generator)
	type testCase[T any] struct {
		name string
		s    *shard[T]
		want bool
	}
	tests := []testCase[int]{
		{"check nil true", &shard[int]{1, 23, gen}, false},
		{"check nil true", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsNil(); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shard_Next(t *testing.T) {
	gen := new(generator)
	type testCase[T any] struct {
		name string
		s    *shard[T]
	}
	tests := []testCase[int]{
		{"check next 1", &shard[int]{1, 23, gen}},
		{"check next 1", &shard[int]{2, 24, gen}},
		{"check next 1", &shard[int]{3, 25, gen}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Next(); got == 0 {
				t.Errorf("Next() = %v", got)
			}
		})
	}
}

func Test_shard_T(t *testing.T) {
	gen := new(generator)
	type testCase[T any] struct {
		name string
		s    *shard[T]
		want T
	}
	tests := []testCase[int]{
		{"check value 1", &shard[int]{1, 23, gen}, 23},
		{"check value 2", &shard[int]{2, 24, gen}, 24},
		{"check value 3", &shard[int]{3, 25, gen}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.T(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("T() = %v, want %v", got, tt.want)
			}
		})
	}
}

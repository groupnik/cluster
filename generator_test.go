package cluster

import "testing"

func Test_generator_generate(t *testing.T) {
	type fields struct {
		counter uint32
	}
	type args struct {
		shardID uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"check generate 1", fields{0}, args{1}},
		{"check generate 2", fields{0}, args{2}},
		{"check generate 3", fields{0}, args{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &generator{
				counter: tt.fields.counter,
			}
			got := g.generate(tt.args.shardID)
			t.Logf("generated id: %d", got)
			if got == 0 {
				t.Errorf("generate() = %v, got > 0", got)
			}
		})
	}
}

func Test_generator_next(t *testing.T) {
	gen := new(generator)

	tests := []struct {
		name string

		want uint64
	}{
		{"check next 1", 1},
		{"check next 2", 2},
		{"check next 3", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gen.next(); got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generator_shardID(t *testing.T) {
	gen := new(generator)

	type args struct {
		id uint64
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"check shard 1", args{6144084783099020289}, 1},
		{"check shard 1", args{6144084783099021313}, 2},
		{"check shard 1", args{6144084783099022337}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gen.shardID(tt.args.id); got != tt.want {
				t.Errorf("shardID() = %v, want %v", got, tt.want)
			}
		})
	}
}

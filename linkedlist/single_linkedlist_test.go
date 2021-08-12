package linkedlist

import "testing"

func TestAddNode(t *testing.T) {
	type args struct {
		n *Node
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nil, 1}, 1},
		{"2", args{nil, 12}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddNode(tt.args.n, tt.args.v); got != tt.want {
				t.Errorf("AddNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

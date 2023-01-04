package main

import (
	"math"
	"reflect"
	"testing"
)

func Test_tail_distanceFromHead(t *testing.T) {
	type args struct {
		head head
	}
	tests := []struct {
		name string
		t    tail
		args args
		want float64
	}{
		{
			name: "head right of tail is 1", args: args{head: head{
				x: 1,
				y: 0,
			}}, t: tail{
				x: 0,
				y: 0,
			}, want: 1,
		},
		{
			name: "head 1 right 1 up of tail is sqrt(2)", args: args{head: head{
				x: 1,
				y: 1,
			}}, t: tail{
				x: 0,
				y: 0,
			}, want: math.Sqrt2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.distanceFromHead(tt.args.head); got != tt.want {
				t.Errorf("distanceFromHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part01(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testinput",
			args: args{input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`}, want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part01(tt.args.input); got != tt.want {
				t.Errorf("part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMoveFromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want move
	}{
		{
			name: "R 4 gets right with 4 steps",
			args: args{line: "R 4"},
			want: move{
				steps:     4,
				direction: "R",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoveFromLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMoveFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveTail(t *testing.T) {
	type args struct {
		t *tail
		h *head
	}
	tests := []struct {
		name string
		args args
		want *tail
	}{
		{
			name: "head moved right", args: args{
				t: &tail{
					x: 0,
					y: 0,
				},
				h: &head{
					x: 2,
					y: 0,
				},
			}, want: &tail{
				x: 1,
				y: 0,
			},
		},
		{
			name: "head moved left", args: args{
				t: &tail{
					x: 2,
					y: 0,
				},
				h: &head{
					x: 0,
					y: 0,
				},
			}, want: &tail{
				x: 1,
				y: 0,
			},
		},
		{
			name: "head moved up", args: args{
				t: &tail{
					x: 0,
					y: 0,
				},
				h: &head{
					x: 0,
					y: 2,
				},
			}, want: &tail{
				x: 0,
				y: 1,
			},
		},
		{
			name: "head moved down", args: args{
				t: &tail{
					x: 0,
					y: 2,
				},
				h: &head{
					x: 0,
					y: 0,
				},
			}, want: &tail{
				x: 0,
				y: 1,
			},
		},
		{
			name: "head moved up and right", args: args{
				t: &tail{
					x: 0,
					y: 0,
				},
				h: &head{
					x: 2,
					y: 2,
				},
			}, want: &tail{
				x: 1,
				y: 1,
			},
		},
		{
			name: "head moved up and left", args: args{
				t: &tail{
					x: 0,
					y: 0,
				},
				h: &head{
					x: 2,
					y: -2,
				},
			}, want: &tail{
				x: 1,
				y: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveTail(tt.args.t, tt.args.h)
			if !reflect.DeepEqual(tt.args.t, tt.want) {
				t.Errorf("moveTail() = %v, want %v", tt.args.t, tt.want)
			}
		})
	}
}

func Test_part02(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tesinput", args: args{input: `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`}, want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part02(tt.args.input); got != tt.want {
				t.Errorf("part02() = %v, want %v", got, tt.want)
			}
		})
	}
}

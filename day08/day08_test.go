package main

import (
	"reflect"
	"testing"
)

func Test_getVisibleTrees(t *testing.T) {
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
			args: args{input: `30373
25512
65332
33549
35390
`},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVisibleTrees(tt.args.input); got != tt.want {
				t.Errorf("getVisibleTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOuterTrees(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "length 5 is 16", args: args{i: 5}, want: 16,
		},
		{
			name: "length 100 is 396", args: args{i: 100}, want: 396,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOuterTrees(tt.args.i); got != tt.want {
				t.Errorf("getOuterTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVerticalLineHeights(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "30373 gets 3,0,3,7,3", args: args{line: "30373"}, want: []int{3, 0, 3, 7, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLineHeights(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLineHeights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHorizontalLineHeights(t *testing.T) {
	type args struct {
		lines []string
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "3 2 6 3 3 gets 3,2,6,3,3 at index 0", args: args{
				lines: []string{"3", "2", "6", "3", "3"},
				index: 0,
			}, want: []int{3, 2, 6, 3, 3},
		},
		{
			name: "30373 25512 65332 33549 35390 gets 0,5,5,3,5 at index 1", args: args{
				lines: []string{"30373", "25512", "65332", "33549", "35390"},
				index: 1,
			}, want: []int{0, 5, 5, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHorizontalLineHeights(tt.args.lines, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHorizontalLineHeights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeIsVisible(t *testing.T) {
	type args struct {
		treeHeight     int
		horizontalLine []int
		verticalLine   []int
		x              int
		y              int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "3x3 grid visible in middle",
			args: args{
				treeHeight:     5,
				horizontalLine: []int{1, 5, 1},
				verticalLine:   []int{1, 5, 1},
				x:              1,
				y:              1,
			},
			want: true,
		},
		{
			name: "3x3 grid not visible in middle",
			args: args{
				treeHeight:     5,
				horizontalLine: []int{5, 2, 5},
				verticalLine:   []int{5, 2, 5},
				x:              1,
				y:              1,
			},
			want: false,
		}, {
			name: "Height 3 At Index 2 2 with lines [3 5 3 5 3] [6 5 3 3 2] is true",
			args: args{
				treeHeight:     3,
				horizontalLine: []int{3, 5, 3, 5, 3},
				verticalLine:   []int{6, 5, 3, 3, 2},
				x:              2,
				y:              2,
			},
			want: false,
		}, {
			name: "Height 3 At Index 3 2 with lines [3 5 3 5 3] [6 5 3 3 2] is true",
			args: args{
				treeHeight:     3,
				horizontalLine: []int{3, 5, 3, 5, 3},
				verticalLine:   []int{6, 5, 3, 3, 2},
				x:              3,
				y:              2,
			},
			want: false,
		}, {
			name: "Height 3 At Index 3 3 with lines [7 1 3 4 9] [6 5 3 3 2] is true",
			args: args{
				treeHeight:     3,
				horizontalLine: []int{7, 1, 3, 4, 9},
				verticalLine:   []int{6, 5, 3, 3, 2},
				x:              3,
				y:              3,
			},
			want: true,
		},
		{
			name: "Height 5 At Index 1 1 with lines [0 5 5 3 5] [2 5 5 1 2] is true",
			args: args{
				treeHeight:     5,
				horizontalLine: []int{0, 5, 5, 3, 5},
				verticalLine:   []int{2, 5, 5, 1, 2},
				x:              1,
				y:              1,
			},
			want: true,
		}, {
			name: "Height 5 At Index 1 2 with lines [0 5 5 3 5] [6 5 3 3 2] is true",
			args: args{
				treeHeight:     5,
				horizontalLine: []int{0, 5, 5, 3, 5},
				verticalLine:   []int{6, 5, 3, 3, 2},
				x:              1,
				y:              2,
			},
			want: true,
		}, {
			name: "Height 1 At Index 3 1 with lines [7 1 3 4 9] [2 5 5 1 2] is true",
			args: args{
				treeHeight:     1,
				horizontalLine: []int{7, 1, 3, 4, 9},
				verticalLine:   []int{2, 5, 5, 1, 2},
				x:              3,
				y:              1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeIsVisible(tt.args.treeHeight, tt.args.horizontalLine, tt.args.verticalLine, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("treeIsVisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

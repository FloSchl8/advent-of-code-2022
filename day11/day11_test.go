package main

import "testing"

func Test_part01(t *testing.T) {
	type args struct {
		monkeys []*monkey
	}
	testMonkeys := []*monkey{{
		items: []int{79, 98},
		operation: func(old int) int {
			return old * 19
		},
		testMod:    23,
		trueThrow:  2,
		falseThrow: 3,
	}, {
		items: []int{54, 65, 75, 74},
		operation: func(old int) int {
			return old + 6
		},
		testMod:    19,
		trueThrow:  2,
		falseThrow: 0,
	}, {
		items: []int{79, 60, 97},
		operation: func(old int) int {
			return old * old
		},
		testMod:    13,
		trueThrow:  1,
		falseThrow: 3,
	}, {
		items: []int{74},
		operation: func(old int) int {
			return old + 3
		},
		testMod:    17,
		trueThrow:  0,
		falseThrow: 1,
	},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testinput",
			args: args{monkeys: testMonkeys},
			want: 10605,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part01(tt.args.monkeys); got != tt.want {
				t.Errorf("part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part02(t *testing.T) {
	type args struct {
		monkeys []*monkey
	}
	testMonkeys := []*monkey{{
		items: []int{79, 98},
		operation: func(old int) int {
			return old * 19
		},
		testMod:    23,
		trueThrow:  2,
		falseThrow: 3,
	}, {
		items: []int{54, 65, 75, 74},
		operation: func(old int) int {
			return old + 6
		},
		testMod:    19,
		trueThrow:  2,
		falseThrow: 0,
	}, {
		items: []int{79, 60, 97},
		operation: func(old int) int {
			return old * old
		},
		testMod:    13,
		trueThrow:  1,
		falseThrow: 3,
	}, {
		items: []int{74},
		operation: func(old int) int {
			return old + 3
		},
		testMod:    17,
		trueThrow:  0,
		falseThrow: 1,
	},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testinput",
			args: args{monkeys: testMonkeys},
			want: 2713310158,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part02(tt.args.monkeys); got != tt.want {
				t.Errorf("part02() = %v, want %v", got, tt.want)
			}
		})
	}
}

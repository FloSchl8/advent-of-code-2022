package main

import "testing"

func Test_sortBackpackItems(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "testinput", args: args{input: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`}, want: 157},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCharPriorityFromAsciiIndex(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "p sould return 16", args: args{i: 'p'}, want: 16},
		{name: "P sould return 42", args: args{i: 'P'}, want: 42},
		{name: "L sould return 38", args: args{i: 'L'}, want: 38},
		{name: "v sould return 22", args: args{i: 'v'}, want: 22},
		{name: "t sould return 20", args: args{i: 't'}, want: 20},
		{name: "s sould return 19", args: args{i: 's'}, want: 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCharPriorityFromAsciiIndex(tt.args.i); got != tt.want {
				t.Errorf("getCharPriorityFromAsciiIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "testinput", args: args{input: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`}, want: 70}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

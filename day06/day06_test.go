package main

import "testing"

func Test_part01(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "testinput 1", args: args{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, want: 7},
		{name: "testinput 2", args: args{input: "bvwbjplbgvbhsrlpgdmjqwftvncz"}, want: 5},
		{name: "testinput 3", args: args{input: "nppdvjthqldpwncqszvftbrmjlhg"}, want: 6},
		{name: "testinput 4", args: args{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, want: 10},
		{name: "testinput 5", args: args{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part01(tt.args.input); got != tt.want {
				t.Errorf("part01() = %v, want %v", got, tt.want)
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
		{name: "testinput 1", args: args{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, want: 19},
		{name: "testinput 2", args: args{input: "bvwbjplbgvbhsrlpgdmjqwftvncz"}, want: 23},
		{name: "testinput 3", args: args{input: "nppdvjthqldpwncqszvftbrmjlhg"}, want: 23},
		{name: "testinput 4", args: args{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, want: 29},
		{name: "testinput 5", args: args{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, want: 26}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part02(tt.args.input); got != tt.want {
				t.Errorf("part02() = %v, want %v", got, tt.want)
			}
		})
	}
}

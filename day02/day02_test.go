package main

import "testing"

func Test_getScore(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "testinput", args: args{input: `A Y
B X
C Z
`}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScore(tt.args.input); got != tt.want {
				t.Errorf("getScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

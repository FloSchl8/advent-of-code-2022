package main

import (
	"reflect"
	"testing"
)

func Test_stackCrates(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name        string
		args        args
		wantTopOnes string
	}{
		{name: "testinput", args: args{input: `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`}, wantTopOnes: "MCD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTopOnes := stackCrates(tt.args.input); gotTopOnes != tt.wantTopOnes {
				t.Errorf("stackCrates() = %v, want %v", gotTopOnes, tt.wantTopOnes)
			}
		})
	}
}

func Test_findSeperatorLine(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name                           string
		args                           args
		wantInstructionsStartAfterLine int
	}{
		{name: "testinput", args: args{lines: []string{"23234", "fwef", "", "werwer"}}, wantInstructionsStartAfterLine: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInstructionsStartAfterLine := findSeperatorLine(tt.args.lines); gotInstructionsStartAfterLine != tt.wantInstructionsStartAfterLine {
				t.Errorf("findSeperatorLine() = %v, want %v", gotInstructionsStartAfterLine, tt.wantInstructionsStartAfterLine)
			}
		})
	}
}

func Test_getNumberOfStacks(t *testing.T) {
	type args struct {
		input string
	}
	test := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testinput",
			args: args{input: " 1   2   3 "},
			want: 3,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfCrates(tt.args.input); got != tt.want {
				t.Errorf("findSeperatorLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOperation(t *testing.T) {
	type args struct {
		operationInput string
	}
	tests := []struct {
		name string
		args args
		want operation
	}{
		{name: "move 1 from 2 to 1", args: args{"move 1 from 2 to 1"}, want: operation{
			numOfCrates: 1,
			pickup:      2,
			dropAt:      1,
		}},
		{name: "move 3 from 1 to 3", args: args{"move 3 from 1 to 3"}, want: operation{
			numOfCrates: 3,
			pickup:      1,
			dropAt:      3,
		}},
		{name: "move 2 from 2 to 1", args: args{"move 2 from 2 to 1"}, want: operation{
			numOfCrates: 2,
			pickup:      2,
			dropAt:      1,
		}},
		{name: "move 1 from 1 to 2", args: args{"move 1 from 1 to 2"}, want: operation{
			numOfCrates: 1,
			pickup:      1,
			dropAt:      2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOperation(tt.args.operationInput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

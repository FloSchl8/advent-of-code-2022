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
			if got := sortBackpackItems(tt.args.input); got != tt.want {
				t.Errorf("sortBackpackItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func Test_contains(t *testing.T) {
	type args struct {
		s   []rune
		str rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "abc contains a", args: args{
			s:   []rune{'a', 'b', 'c'},
			str: 'a',
		}, want: true},
		{name: "abc contains not d", args: args{
			s:   []rune{'a', 'b', 'c'},
			str: 'd',
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.str); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

package main

import "testing"

func Test_findFullyContainedRanges(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "testinput", args: args{input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFullyContainedRanges(tt.args.input); got != tt.want {
				t.Errorf("findFullyContainedRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitPairInBounds(t *testing.T) {
	type args struct {
		rangeInput string
	}
	tests := []struct {
		name      string
		args      args
		wantLower int
		wantUpper int
	}{
		{name: "2-3 get 2 and 3", args: args{rangeInput: "2-3"}, wantLower: 2, wantUpper: 3},
		{name: "4-5 gets 4 and 5", args: args{rangeInput: "4-5"}, wantLower: 4, wantUpper: 5},
		{name: "5-7 gets 5 and 7", args: args{rangeInput: "5-7"}, wantLower: 5, wantUpper: 7},
		{name: "7-9 gets 7 and 9", args: args{rangeInput: "7-9"}, wantLower: 7, wantUpper: 9},
		{name: "2-8 gets 2 and 8", args: args{rangeInput: "2-8"}, wantLower: 2, wantUpper: 8},
		{name: "3-7 gets 3 and 7", args: args{rangeInput: "3-7"}, wantLower: 3, wantUpper: 7},
		{name: "6-6 gets 6 and 6", args: args{rangeInput: "6-6"}, wantLower: 6, wantUpper: 6},
		{name: "4-6 gets 4 and 6", args: args{rangeInput: "4-6"}, wantLower: 4, wantUpper: 6},
		{name: "2-6 gets 2 and 6", args: args{rangeInput: "2-6"}, wantLower: 2, wantUpper: 6},
		{name: "4-8 gets 4 and 8", args: args{rangeInput: "4-8"}, wantLower: 4, wantUpper: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLower, gotUpper := splitPairInBounds(tt.args.rangeInput)
			if gotLower != tt.wantLower {
				t.Errorf("splitPairInBounds() gotLower = %v, want %v", gotLower, tt.wantLower)
			}
			if gotUpper != tt.wantUpper {
				t.Errorf("splitPairInBounds() gotUpper = %v, want %v", gotUpper, tt.wantUpper)
			}
		})
	}
}

func Test_section_isContained(t *testing.T) {
	type fields struct {
		lowerEnd int
		upperEnd int
	}
	type args struct {
		comparedSection section
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "4-4 is contained in 4-6", fields: fields{
			lowerEnd: 4,
			upperEnd: 4,
		}, args: args{comparedSection: section{
			lowerEnd: 4,
			upperEnd: 6,
		}}, want: true},
		{name: "4-6 is contained in 4-4", fields: fields{
			lowerEnd: 4,
			upperEnd: 6,
		}, args: args{comparedSection: section{
			lowerEnd: 4,
			upperEnd: 4,
		}}, want: true},
		{name: "4-4 is not contained in 5-6", fields: fields{
			lowerEnd: 4,
			upperEnd: 4,
		}, args: args{comparedSection: section{
			lowerEnd: 5,
			upperEnd: 6,
		}}, want: false},
		{name: "4-8 is not contained in 15-18", fields: fields{
			lowerEnd: 4,
			upperEnd: 8,
		}, args: args{comparedSection: section{
			lowerEnd: 15,
			upperEnd: 18,
		}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := section{
				lowerEnd: tt.fields.lowerEnd,
				upperEnd: tt.fields.upperEnd,
			}
			if got := s.isContained(tt.args.comparedSection); got != tt.want {
				t.Errorf("isContained() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_section_isOverlapping(t *testing.T) {
	type fields struct {
		lowerEnd int
		upperEnd int
	}
	type args struct {
		comparedSection section
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "3-4 is overlapping in 4-6", fields: fields{
			lowerEnd: 3,
			upperEnd: 4,
		}, args: args{comparedSection: section{
			lowerEnd: 4,
			upperEnd: 6,
		}}, want: true},
		{name: "4-6 is overlapping in 6-8", fields: fields{
			lowerEnd: 4,
			upperEnd: 6,
		}, args: args{comparedSection: section{
			lowerEnd: 6,
			upperEnd: 8,
		}}, want: true},
		{name: "4-8 is overlapping in 5-9", fields: fields{
			lowerEnd: 4,
			upperEnd: 8,
		}, args: args{comparedSection: section{
			lowerEnd: 5,
			upperEnd: 9,
		}}, want: true},
		{name: "4-8 is not overlapping in 15-18", fields: fields{
			lowerEnd: 4,
			upperEnd: 8,
		}, args: args{comparedSection: section{
			lowerEnd: 15,
			upperEnd: 18,
		}}, want: false},
		{name: "6-6 is not overlapping in 4-6", fields: fields{
			lowerEnd: 6,
			upperEnd: 6,
		}, args: args{comparedSection: section{
			lowerEnd: 4,
			upperEnd: 6,
		}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := section{
				lowerEnd: tt.fields.lowerEnd,
				upperEnd: tt.fields.upperEnd,
			}
			if got := s.isOverlapping(tt.args.comparedSection); got != tt.want {
				t.Errorf("isOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOverallpingRanges(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "testinput", args: args{input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
83-85,8-84
`}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOverallpingRanges(tt.args.input); got != tt.want {
				t.Errorf("findOverallpingRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

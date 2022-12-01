package main

import (
	"reflect"
	"testing"
)

func Test_sumCaolries(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		{name: "testinput", args: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`, want: 45000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumCaolries(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumCaolries() = %v, want %v", got, tt.want)
			}
		})
	}
}

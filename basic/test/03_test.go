package test

import (
	"fmt"
	"testing"
)

func Test_myAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args)
			if got := myAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("myAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

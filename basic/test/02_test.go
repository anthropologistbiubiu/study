package test

import "testing"

func Test_containerTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"sun"},
		{"wei"},
		{"ming"}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			containerTest()
		})
	}
}

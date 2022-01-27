package main

import "testing"

func TestUnlimitChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"测试永动机"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UnlimitChan()
		})
	}
}

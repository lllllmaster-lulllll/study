package test

import (
	"testing"
)

func TestFallThrough(t *testing.T) {

	FallThrough()
}

func Test_sliceCut(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceCut()
		})
	}
}

package structures

import (
	"fmt"
	"testing"
)

func TestSpace_IsIncludes(t *testing.T) {
	tests := []struct {
		space Volume
		dot   Dot
		want  bool
	}{
		{
			space: Volume{Start: Dot{-10, -10, -10}, End: Dot{10, 10, 10}},
			dot:   Dot{0, 0, 0},
			want:  true,
		},
		{
			space: EmptySpace(),
			dot:   Dot{0, 0, 0},
			want:  false,
		},
		{
			space: Volume{Start: Dot{-10, -10, -10}, End: Dot{10, 10, 10}},
			dot:   Dot{20, 0, 0},
			want:  false,
		},
		{
			space: Volume{Start: Dot{-10, -10, -10}, End: Dot{10, 10, 10}},
			dot:   Dot{0, -20, 0},
			want:  false,
		},
		{
			space: Volume{Start: Dot{-10, -10, -10}, End: Dot{10, 10, 10}},
			dot:   Dot{0, 0, -20},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("IsIncludes %v in %v", tt.dot, tt.space), func(t *testing.T) {
			if got := tt.space.IsIncludes(tt.dot); got != tt.want {
				t.Errorf("IsIncludes() = %v, want %v", got, tt.want)
			}
		})
	}
}

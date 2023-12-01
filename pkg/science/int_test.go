package science

import (
	"strconv"
	"testing"
)

func TestAbsInt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			x:    2,
			want: 2,
		},
		{
			x:    -5,
			want: 5,
		},
		{
			x:    0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.x), func(t *testing.T) {
			if got := AbsInt(tt.x); got != tt.want {
				t.Errorf("AbsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumElementsInt(t *testing.T) {
	tests := []struct {
		elements []int
		want     int
	}{
		{
			elements: []int{199, 200, 208},
			want:     607,
		},
		{
			elements: []int{155, -200},
			want:     -45,
		},
		{
			elements: []int{},
			want:     0,
		},
	}
	for i, tt := range tests {
		t.Run("Test "+strconv.Itoa(i), func(t *testing.T) {
			if got := SumElementsInt(tt.elements); got != tt.want {
				t.Errorf("SumElementsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

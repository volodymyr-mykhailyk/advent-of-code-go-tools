package science

import (
	"fmt"
	"testing"
)

func TestLeastCommonMultiple(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{
			numbers: []int{3},
			want:    3,
		},
		{
			numbers: []int{10, 15, 20},
			want:    60,
		},
		{
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:    2520,
		},
		{
			numbers: []int{16343, 13019, 21883, 20221, 11911, 19667},
			want:    13524038372771,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.numbers), func(t *testing.T) {
			if got := LeastCommonMultiple(tt.numbers); got != tt.want {
				t.Errorf("LeastCommonMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreatCommonDivisor(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			a:    3,
			b:    6,
			want: 3,
		},
		{
			a:    7,
			b:    21,
			want: 7,
		},
		{
			a:    1260,
			b:    210,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v,%v", tt.a, tt.b), func(t *testing.T) {
			if got := GreatCommonDivisor(tt.a, tt.b); got != tt.want {
				t.Errorf("GreatCommonDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

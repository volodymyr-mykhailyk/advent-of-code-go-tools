package structures

import (
	"sort"
)

type Volume struct {
	Start Dot
	End   Dot
}

func EmptySpace() Volume {
	return Volume{Start: Dot{1, 1, 1}, End: Dot{-1, -1, -1}}
}

func (s Volume) VisitAll(f func(dot Dot)) {
	xs, ys, zs := s.Start.Values()
	xe, ye, ze := s.End.Values()
	for x := xs; x <= xe; x++ {
		for y := ys; y <= ye; y++ {
			for z := zs; z <= ze; z++ {
				f(Dot{X: x, Y: y, Z: z})
			}
		}
	}
}

func (s Volume) CornerDots() []Dot {
	if s.IsEmpty() {
		return []Dot{}
	}

	corners := make([]Dot, 8)
	xs, ys, zs := s.Start.Values()
	xe, ye, ze := s.End.Values()
	i := 0
	for _, x := range []int{xs, xe} {
		for _, y := range []int{ys, ye} {
			for _, z := range []int{zs, ze} {
				corners[i] = Dot{X: x, Y: y, Z: z}
				i++
			}
		}
	}
	return corners
}

func (s Volume) IntersectionWith(s2 Volume) Volume {
	if s.IsIntersectsWith(s2) {
		xs, xe := intersectionRange(s.Start.X, s.End.X, s2.Start.X, s2.End.X)
		ys, ye := intersectionRange(s.Start.Y, s.End.Y, s2.Start.Y, s2.End.Y)
		zs, ze := intersectionRange(s.Start.Z, s.End.Z, s2.Start.Z, s2.End.Z)
		return Volume{Start: Dot{xs, ys, zs}, End: Dot{xe, ye, ze}}
	} else {
		return EmptySpace()
	}
}

func (s Volume) IsIntersectsWith(s2 Volume) bool {
	for _, corner := range s.CornerDots() {
		if s2.IsIncludes(corner) {
			return true
		}
	}
	return false
}

func (s Volume) IsEmpty() bool {
	if s == EmptySpace() {
		return true
	} else {
		return false
	}
}

func (s Volume) IsIncludes(dot Dot) bool {
	xs, ys, zs := s.Start.Values()
	xe, ye, ze := s.End.Values()
	includeX := xs <= dot.X && dot.X <= xe
	includeY := ys <= dot.Y && dot.Y <= ye
	includeZ := zs <= dot.Z && dot.Z <= ze
	if includeX && includeY && includeZ {
		return true
	} else {
		return false
	}
}

func (s Volume) Size() int {
	if s.IsEmpty() {
		return 0
	} else {
		xs, ys, zs := s.Start.Values()
		xe, ye, ze := s.End.Values()
		return (xe - xs) * (ye - ys) * (ze - zs)
	}
}

func intersectionRange(x1 int, x2 int, x3 int, x4 int) (int, int) {
	sorted := []int{x1, x2, x3, x4}
	sort.Ints(sorted)
	return sorted[1], sorted[2]
}

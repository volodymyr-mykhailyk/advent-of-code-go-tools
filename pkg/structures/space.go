package structures

import (
	"math"
)

type Space [][][]int

func (s Space) Get(dot Dot) int {
	if s.IsInbound(dot) {
		return s[dot.Z][dot.Y][dot.X]
	} else {
		return math.MaxInt
	}
}

func (s Space) Set(dot Dot, value int) {
	if s.IsInbound(dot) {
		s[dot.Z][dot.Y][dot.X] = value
	}
}

func (s Space) VisitAll(f func(value int, dot Dot)) {
	for z, plane := range s {
		for y, row := range plane {
			for x, value := range row {
				f(value, Dot{X: x, Y: y, Z: z})
			}
		}
	}
}

func (s Space) VisitSurrounding(dot Dot, f func(value int, dot Dot)) {
	dot.VisitSurroundings(func(surrounding Dot) {
		s.Visit(surrounding, f)
	})
}

func (s Space) Visit(dot Dot, f func(value int, dot Dot)) {
	if s.IsInbound(dot) {
		f(s.Get(dot), dot)
	}
}

func (s Space) IsInbound(dot Dot) bool {
	x, y, z := dot.Coords()
	if z < 0 || z >= len(s) {
		return false
	} else if y < 0 || y >= len(s[z]) {
		return false
	} else if x < 0 || x >= len(s[z][y]) {
		return false
	} else {
		return true
	}
}

func (s Space) Size() int {
	return len(s) * len(s[0]) * len(s[0][0])
}

func (s Space) Dimensions() (int, int, int) {
	return len(s[0][0]), len(s[0]), len(s)
}

func (s Space) Start() Dot {
	return Dot{X: 0, Y: 0, Z: 0}
}

func (s Space) End() Dot {
	x, y, z := s.Dimensions()
	return Dot{X: x - 1, Y: y - 1, Z: z - 1}
}

func (s Space) IsAtEdge(dot Dot) bool {
	end := s.End()
	if dot.X == 0 || dot.Y == 0 || dot.Z == 0 {
		return true
	}
	if dot.X == end.X || dot.Y == end.Y || dot.Z == end.Z {
		return true
	}
	return false
}

func SpaceFromSpec(maxX, maxY, maxZ, value int) Space {
	space := make(Space, maxZ)
	for z := 0; z < maxZ; z++ {
		space[z] = make([][]int, maxY)
		for y := 0; y < maxY; y++ {
			space[z][y] = make([]int, maxX)
			for x := range space[z][y] {
				space[z][y][x] = value
			}
		}
	}
	return space
}

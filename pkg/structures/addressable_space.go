package structures

import (
	"github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/science"
	"math"
)

type AddressableSpace map[Dot]int

func (space AddressableSpace) Get(dot Dot) int {
	if value, ok := space[dot]; ok {
		return value
	} else {
		return 0
	}
}

func (space AddressableSpace) Set(dot Dot, value int) {
	space[dot] = value
}

func (space AddressableSpace) VisitAll(f func(dot Dot, value int)) {
	for dot, value := range space {
		f(dot, value)
	}
}

func (space AddressableSpace) End() Dot {
	maxX, maxY, maxZ := math.MinInt, math.MinInt, math.MinInt
	space.VisitAll(func(dot Dot, _ int) {
		maxX = science.MaxFromPairInt(maxX, dot.X)
		maxY = science.MaxFromPairInt(maxY, dot.Y)
		maxZ = science.MaxFromPairInt(maxZ, dot.Z)
	})
	return Dot{X: maxX, Y: maxY, Z: maxZ}
}

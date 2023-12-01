package structures

import (
	"github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/transformations"
	"strings"
)

type Dot struct {
	X int
	Y int
	Z int
}

func (d Dot) Values() (int, int, int) {
	return d.X, d.Y, d.Z
}

func (d Dot) VisitSurroundings(f func(dot Dot)) {
	f(d.Up())
	f(d.Right())
	f(d.Down())
	f(d.Left())
	f(d.Front())
	f(d.Back())
}

func (d Dot) Up() Dot {
	return Dot{X: d.X, Y: d.Y - 1, Z: d.Z}
}

func (d Dot) Right() Dot {
	return Dot{X: d.X + 1, Y: d.Y, Z: d.Z}
}

func (d Dot) Down() Dot {
	return Dot{X: d.X, Y: d.Y + 1, Z: d.Z}
}

func (d Dot) Left() Dot {
	return Dot{X: d.X - 1, Y: d.Y, Z: d.Z}
}

func (d Dot) Front() Dot {
	return Dot{X: d.X, Y: d.Y, Z: d.Z + 1}
}

func (d Dot) Back() Dot {
	return Dot{X: d.X, Y: d.Y, Z: d.Z - 1}
}

func (d Dot) Coords() (int, int, int) {
	return d.X, d.Y, d.Z
}

func DotFromString(input string) Dot {
	coordinates := transformations.ParseIntegers(strings.Split(input, ","))
	return Dot{X: coordinates[0], Y: coordinates[1], Z: coordinates[2]}
}

package structures

import (
	"fmt"
	"github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/science"
	"github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/transformations"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (point Point) Up() Point {
	return Point{X: point.X, Y: point.Y - 1}
}

func (point Point) UpRight() Point {
	return Point{X: point.X + 1, Y: point.Y - 1}
}

func (point Point) Right() Point {
	return Point{X: point.X + 1, Y: point.Y}
}

func (point Point) DownRight() Point {
	return Point{X: point.X + 1, Y: point.Y + 1}
}

func (point Point) Down() Point {
	return Point{X: point.X, Y: point.Y + 1}
}

func (point Point) DownLeft() Point {
	return Point{X: point.X - 1, Y: point.Y + 1}
}

func (point Point) Left() Point {
	return Point{X: point.X - 1, Y: point.Y}
}

func (point Point) UpLeft() Point {
	return Point{X: point.X - 1, Y: point.Y - 1}
}

func (point Point) Coords() (int, int) {
	return point.X, point.Y
}

func (point Point) ByDirection(direction string) Point {
	switch direction {
	case "U":
		return point.Up()
	case "R":
		return point.Right()
	case "D":
		return point.Down()
	case "L":
		return point.Left()
	default:
		panic("Unknown move")
	}
}

func (point Point) String() string {
	return fmt.Sprintf("{%v,%v}", point.X, point.Y)
}

func (point Point) ManhattanDistanceTo(target Point) int {
	return science.AbsInt(point.X-target.X) + science.AbsInt(point.Y-target.Y)
}

func PointFromString(input string) Point {
	coordinates := transformations.ParseIntegers(strings.Split(input, ","))
	return Point{X: int(coordinates[0]), Y: int(coordinates[1])}
}

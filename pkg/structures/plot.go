package structures

import (
	"github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/transformations"
	"math"
	"strconv"
	"strings"
)

type Plot [][]int

func (plot Plot) Get(point Point) int {
	if plot.IsInbound(point) {
		return plot[point.Y][point.X]
	} else {
		return math.MaxInt
	}
}

func (plot Plot) Set(point Point, value int) {
	if plot.IsInbound(point) {
		plot[point.Y][point.X] = value
	}
}

func (plot Plot) Surroundings(point Point) []int {
	results := make([]int, 4)
	results[0] = plot.Get(point.Up())
	results[1] = plot.Get(point.Right())
	results[2] = plot.Get(point.Down())
	results[3] = plot.Get(point.Left())
	return results
}

func (plot Plot) VisitAll(f func(value int, point Point)) {
	for y, row := range plot {
		for x, value := range row {
			f(value, Point{X: x, Y: y})
		}
	}
}

func (plot Plot) VisitAdjacent(point Point, f func(value int, point Point)) {
	plot.VisitPoint(point.Up(), f)
	plot.VisitPoint(point.Right(), f)
	plot.VisitPoint(point.Down(), f)
	plot.VisitPoint(point.Left(), f)
}

func (plot Plot) VisitAround(point Point, f func(value int, point Point)) {
	plot.VisitPoint(point.Up(), f)
	plot.VisitPoint(point.UpRight(), f)
	plot.VisitPoint(point.Right(), f)
	plot.VisitPoint(point.DownRight(), f)
	plot.VisitPoint(point.Down(), f)
	plot.VisitPoint(point.DownLeft(), f)
	plot.VisitPoint(point.Left(), f)
	plot.VisitPoint(point.UpLeft(), f)
}

func (plot Plot) VisitPoint(point Point, f func(value int, point Point)) {
	if plot.IsInbound(point) {
		f(plot.Get(point), point)
	}
}

func (plot Plot) IsInbound(point Point) bool {
	x, y := point.X, point.Y
	if y < 0 || y >= len(plot) {
		return false
	} else if x < 0 || x >= len(plot[y]) {
		return false
	} else {
		return true
	}
}

func (plot Plot) Refresh(f func(v int, point Point) int) {
	plot.VisitAll(func(v int, point Point) {
		plot.Set(point, f(v, point))
	})
}

func (plot Plot) VLine(x int) []int {
	slice := make([]int, len(plot))
	for y, line := range plot {
		slice[y] = line[x]
	}
	return slice
}

func (plot Plot) HLine(y int) []int {
	return plot[y]
}

func (plot Plot) Presentation(valueSep string) string {
	result := make([]string, len(plot))
	for y, line := range plot {
		elements := make([]string, len(line))
		for x, value := range line {
			elements[x] = strconv.FormatInt(int64(value), 10)
		}
		result[y] = strings.Join(elements, valueSep)
	}
	return strings.Join(result, "\n")
}

func (plot Plot) Sprites(valueSep string, sprites map[int]string) string {
	result := make([]string, len(plot))
	for y, line := range plot {
		elements := make([]string, len(line))
		for x, value := range line {
			elements[x] = sprites[value]
		}
		result[y] = strings.Join(elements, valueSep)
	}
	return strings.Join(result, "\n")
}

func (plot Plot) Size() int {
	return len(plot) * len(plot[0])
}

func (plot Plot) Dimensions() (int, int) {
	return len(plot[0]), len(plot)
}

func (plot Plot) Start() Point {
	return Point{X: 0, Y: 0}
}

func (plot Plot) End() Point {
	return Point{X: len(plot[0]) - 1, Y: len(plot) - 1}
}

func (plot Plot) Width() int {
	return len(plot[0])
}

func (plot Plot) Height() int {
	return len(plot)
}

func (plot Plot) SurroundingPoints(point Point) []Point {
	return []Point{
		point.Up(),
		point.Right(),
		point.Down(),
		point.Left(),
	}
}

func PlotFromString(input [][]string) Plot {
	maxY := len(input)
	maxX := len(input[0])
	plot := make(Plot, maxY)
	for y, line := range input {
		plot[y] = make([]int, maxX)
		for x, value := range line {
			plot[y][x] = transformations.ParseInteger(value)
		}
	}
	return plot
}

func PlotFromSpec(maxX, maxY, value int) Plot {
	plot := make(Plot, maxY)
	for y := 0; y < maxY; y++ {
		plot[y] = make([]int, maxX)
		for x := range plot[y] {
			plot[y][x] = value
		}
	}
	return plot
}

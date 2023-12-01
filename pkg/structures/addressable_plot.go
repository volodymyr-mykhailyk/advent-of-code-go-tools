package structures

import (
	"math"
	"strconv"
	"strings"
)

type AddressablePlot map[Point]int

func (plot AddressablePlot) Presentation() string {
	minX, maxX := plot.HDimensions()
	minY, maxY := plot.VDimensions()
	presentation := make([]string, maxY-minY+1)
	for y := range presentation {
		presentation[y] = strings.Repeat(".", maxX-minX+1)
	}
	for point, value := range plot {
		if value != math.MinInt {
			x := point.X - minX
			y := point.Y - minY
			s := presentation[y]
			presentation[y] = s[:x] + strconv.Itoa(value) + s[x+1:]
		}
	}
	return strings.Join(presentation, "\n")
}

func (plot AddressablePlot) SpritePresentation(sprites map[int]string) string {
	minX, maxX := plot.HDimensions()
	minY, maxY := plot.VDimensions()
	presentation := make([]string, maxY-minY+1)
	for y := range presentation {
		presentation[y] = strings.Repeat(sprites[0], maxX-minX+1)
	}
	for point, value := range plot {
		if value != math.MaxInt {
			x := point.X - minX
			y := point.Y - minY
			s := presentation[y]
			presentation[y] = s[:x] + sprites[value] + s[x+1:]
		}
	}
	return strings.Join(presentation, "\n")
}

func (plot AddressablePlot) Size() int {
	minX, maxX := plot.HDimensions()
	minY, maxY := plot.VDimensions()
	return ((maxX - minX) + 1) * ((maxY - minY) + 1)
}

func (plot AddressablePlot) HDimensions() (int, int) {
	min := math.MaxInt
	max := math.MinInt
	for point := range plot {
		if max < point.X {
			max = point.X
		}
		if min > point.X {
			min = point.X
		}
	}
	return min, max
}

func (plot AddressablePlot) VDimensions() (int, int) {
	min := math.MaxInt
	max := math.MinInt
	for point := range plot {
		if max < point.Y {
			max = point.Y
		}
		if min > point.Y {
			min = point.Y
		}
	}
	return min, max
}

func (plot AddressablePlot) Get(point Point) int {
	if value, ok := plot[point]; ok {
		return value
	} else {
		return 0
	}
}

func (plot AddressablePlot) Set(point Point, value int) {
	plot[point] = value
}

func (plot AddressablePlot) Delete(point Point) {
	delete(plot, point)
}

func (plot AddressablePlot) Copy() AddressablePlot {
	newPlot := make(AddressablePlot)
	for point, v := range plot {
		newPlot[point] = v
	}
	return newPlot
}

func (plot AddressablePlot) VisitAll(f func(point Point, value int)) {
	for point, v := range plot {
		f(point, v)
	}
}

func (plot AddressablePlot) VisitAround(point Point, f func(point Point, value int)) {
	plot.VisitPoint(point.Up(), f)
	plot.VisitPoint(point.UpRight(), f)
	plot.VisitPoint(point.Right(), f)
	plot.VisitPoint(point.DownRight(), f)
	plot.VisitPoint(point.Down(), f)
	plot.VisitPoint(point.DownLeft(), f)
	plot.VisitPoint(point.Left(), f)
	plot.VisitPoint(point.UpLeft(), f)
}

func (plot AddressablePlot) VisitPoint(point Point, f func(point Point, value int)) {
	f(point, plot.Get(point))
}

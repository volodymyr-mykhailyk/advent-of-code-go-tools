package structures

import "github.com/volodymyr-mykhailyk/advent-of-code-go-tools/pkg/science"

type PlotLine struct {
	Start Point
	End   Point
}

func (l PlotLine) MinX() int {
	return science.MinFromPairInt(l.Start.X, l.End.X)
}

func (l PlotLine) MaxX() int {
	return science.MaxFromPairInt(l.Start.X, l.End.X)
}

func (l PlotLine) VisitAll(f func(point Point)) {
	if l.IsHorizontal() {
		for i := l.MinX(); i <= l.MaxX(); i++ {
			f(Point{X: i, Y: l.Start.Y})
		}
	} else if l.IsVertical() {
		for i := l.MinY(); i <= l.MaxY(); i++ {
			f(Point{X: l.Start.X, Y: i})
		}
	}
}

func (l PlotLine) IsHorizontal() bool {
	return l.Start.Y == l.End.Y
}

func (l PlotLine) IsVertical() bool {
	return l.Start.X == l.End.X
}

func (l PlotLine) MinY() int {
	return science.MinFromPairInt(l.Start.Y, l.End.Y)
}

func (l PlotLine) MaxY() int {
	return science.MaxFromPairInt(l.Start.Y, l.End.Y)
}

package structures

type Area struct {
	Xs, Xe int
	Ys, Ye int
}

func (area Area) Includes(point Point) bool {
	if point.X < area.Xs || point.X > area.Xe {
		return false
	}
	if point.Y > area.Ys || point.Y < area.Ye {
		return false
	}

	return true
}

func (area Area) Width() int {
	return area.Xe - area.Xs
}

func (area Area) Height() int {
	return area.Ys - area.Ye
}

func (area Area) IsToTheLeftFrom(point Point) bool {
	if point.X > area.Xe {
		return true
	}
	return false
}

func (area Area) IsHigherFrom(point Point) bool {
	if point.Y < area.Ye {
		return true
	}
	return false
}

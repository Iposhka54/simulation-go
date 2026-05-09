package _map

type Coordinates struct {
	X int
	Y int
}

func NewCoordinates(x, y int) Coordinates {
	return Coordinates{
		X: x,
		Y: y,
	}
}

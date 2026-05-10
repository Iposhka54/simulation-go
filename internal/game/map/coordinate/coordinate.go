package coordinate

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func New(x, y int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("Coords{X=%d, Y=%d}", c.X, c.Y)
}

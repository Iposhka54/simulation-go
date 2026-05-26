package coordinate

import "fmt"

type Point struct {
	X int
	Y int
}

func New(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (c *Point) String() string {
	return fmt.Sprintf("Point{X=%d, Y=%d}", c.X, c.Y)
}

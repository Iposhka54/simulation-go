package static

import (
	"simulation/internal/game/map/coordinate"
)

type Rock struct {
	*StaticEntity
}

func NewRock(c coordinate.Coordinate) *Rock {
	return &Rock{
		StaticEntity: NewStaticEntity(c),
	}
}

func (r *Rock) String() string {
	//todo: move constant in utility class
	return "rock"
}

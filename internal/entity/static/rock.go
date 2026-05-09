package static

import (
	"simulation/internal/entity"
	"simulation/internal/game/map/coordinate"
)

type Rock struct {
	*entity.BaseEntity
}

func NewRock(c coordinate.Coordinate) *Rock {
	return &Rock{
		BaseEntity: entity.New(c),
	}
}

func (r *Rock) String() string {
	//todo: move constant in utility class
	return "rock"
}

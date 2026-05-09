package static

import (
	"simulation/internal/entity"
	"simulation/internal/game/map/coordinate"
)

type Grass struct {
	*entity.BaseEntity
}

func NewGrass(c coordinate.Coordinate) *Grass {
	return &Grass{
		BaseEntity: entity.New(c),
	}
}

func (g *Grass) String() string {
	//todo: move constant in utility class
	return "grass"
}

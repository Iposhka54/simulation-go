package static

import (
	"simulation/internal/game/map/coordinate"
)

type Grass struct {
	*StaticEntity
}

func NewGrass(c coordinate.Coordinate) *Grass {
	return &Grass{
		StaticEntity: NewStaticEntity(c),
	}
}

func (g *Grass) String() string {
	//todo: move constant in utility class
	return "grass"
}

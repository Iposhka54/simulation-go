package static

import (
	"simulation/internal/game/map/coordinate"
)

type Tree struct {
	*StaticEntity
}

func NewTree(c coordinate.Coordinate) *Tree {
	return &Tree{
		StaticEntity: NewStaticEntity(c),
	}
}

func (t *Tree) String() string {
	//todo: move constant in utility class
	return "tree"
}

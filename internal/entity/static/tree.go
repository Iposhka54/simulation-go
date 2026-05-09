package static

import (
	"simulation/internal/entity"
	"simulation/internal/game/map/coordinate"
)

type Tree struct {
	*entity.BaseEntity
}

func NewTree(c coordinate.Coordinate) *Tree {
	return &Tree{
		BaseEntity: entity.New(c),
	}
}

func (t *Tree) String() string {
	//todo: move constant in utility class
	return "tree"
}

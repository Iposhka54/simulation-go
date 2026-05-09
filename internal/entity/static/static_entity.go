package static

import (
	"simulation/internal/entity"
	"simulation/internal/game/map/coordinate"
)

type StaticEntity struct {
	*entity.BaseEntity
}

func NewStaticEntity(c coordinate.Coordinate) *StaticEntity {
	return &StaticEntity{
		BaseEntity: entity.New(c),
	}
}

func (se *StaticEntity) MakeMove(to coordinate.Coordinate) {
	//nothing, because static entity
}

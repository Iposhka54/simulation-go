package static

import (
	"simulation/internal/entity"
)

type StaticEntity struct {
	*entity.BaseEntity
}

func NewStaticEntity() *StaticEntity {
	return &StaticEntity{
		BaseEntity: entity.New(),
	}
}

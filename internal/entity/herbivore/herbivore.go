package herbivore

import (
	"simulation/internal/entity"
	"simulation/internal/entity/creature"
	_map "simulation/internal/game/map"
)

type Herbivore struct {
	*creature.BaseCreature
}

func New(hp, maxHp, speed int) *Herbivore {
	return &Herbivore{
		BaseCreature: creature.New(hp, maxHp, speed),
	}
}

func (h *Herbivore) MakeMove(m *_map.Map) {
	_ = m
	//todo: 1-step randomly move
}

func (h *Herbivore) Type() entity.EntityType {
	return entity.TypeHerbivore
}

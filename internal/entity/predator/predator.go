package predator

import (
	"simulation/internal/entity"
	"simulation/internal/entity/creature"
	_map "simulation/internal/game/map"
)

type Predator struct {
	*creature.BaseCreature
	damage int
}

func New(hp, maxHp, speed, damage int) *Predator {
	return &Predator{
		BaseCreature: creature.New(hp, maxHp, speed),
		damage:       damage,
	}
}

func (h *Predator) MakeMove(m *_map.Map) {
	_ = m
	//todo: 1-step randomly move
}

func (h *Predator) Type() entity.EntityType {
	return entity.TypePredator
}

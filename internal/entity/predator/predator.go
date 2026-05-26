package predator

import (
	"simulation/internal/entity"
	"simulation/internal/entity/creature"
	"simulation/internal/entity/herbivore"
	_map "simulation/internal/game/map"
	"simulation/internal/game/map/coordinate"
	"simulation/internal/game/path"
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

func (p *Predator) MakeMove(m *_map.Map) {
	p.BaseCreature.PerformMove(p, m)
}

func (p *Predator) HasAdjacentFood(m *_map.Map) bool {
	_, _, exists := p.findAdjacentFood(m)
	return exists
}

func (p *Predator) EatAdjacentFood(m *_map.Map) bool {
	foodPosition, prey, exists := p.findAdjacentFood(m)
	if !exists {
		return false
	}

	prey.TakeDamage(p.damage)
	if !prey.IsAlive() {
		m.RemoveEntity(foodPosition)
	}

	return true
}

func (p *Predator) IsFoodAdjacent(m *_map.Map, c coordinate.Coordinate) bool {
	for _, neighbor := range path.GetNeighbors(c) {
		if !m.IsValid(neighbor) {
			continue
		}

		e := m.Get(neighbor.X, neighbor.Y)
		if e == nil {
			continue
		}

		if isHerbivore(e) {
			return true
		}
	}

	return false
}

func (p *Predator) findAdjacentFood(m *_map.Map) (coordinate.Coordinate, creature.Creature, bool) {
	currentPosition := m.GetCoordinatesByEntity(p)
	for _, neighbor := range path.GetNeighbors(currentPosition) {
		if !m.IsValid(neighbor) {
			continue
		}

		e := m.Get(neighbor.X, neighbor.Y)
		if e == nil || !isHerbivore(e) {
			continue
		}

		prey, ok := e.(creature.Creature)
		if !ok {
			continue
		}

		return neighbor, prey, true
	}

	return coordinate.Coordinate{}, nil, false
}

func isHerbivore(e entity.Entity) bool {
	_, ok := e.(*herbivore.Rabbit)
	return ok
}

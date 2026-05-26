package predator

import (
	"simulation/internal/entity"
	"simulation/internal/entity/creature"
	"simulation/internal/entity/herbivore"
	"simulation/internal/game/path"
	"simulation/internal/game/world"
	"simulation/internal/game/world/coordinate"
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

func (p *Predator) MakeMove(w *world.World) {
	p.BaseCreature.PerformMove(p, w)
}

func (p *Predator) HasAdjacentFood(w *world.World) bool {
	_, _, exists := p.findAdjacentFood(w)
	return exists
}

func (p *Predator) EatAdjacentFood(w *world.World) bool {
	foodPosition, prey, exists := p.findAdjacentFood(w)
	if !exists {
		return false
	}

	prey.TakeDamage(p.damage)
	if !prey.IsAlive() {
		w.RemoveEntity(foodPosition)
	}

	return true
}

func (p *Predator) IsFoodAdjacent(w *world.World, point coordinate.Point) bool {
	for _, neighbor := range path.GetNeighbors(point) {
		if !w.IsValid(neighbor) {
			continue
		}

		e := w.Get(neighbor.X, neighbor.Y)
		if e == nil {
			continue
		}

		if isHerbivore(e) {
			return true
		}
	}

	return false
}

func (p *Predator) findAdjacentFood(w *world.World) (coordinate.Point, creature.Creature, bool) {
	currentPosition := w.GetPointByEntity(p)
	for _, neighbor := range path.GetNeighbors(currentPosition) {
		if !w.IsValid(neighbor) {
			continue
		}

		e := w.Get(neighbor.X, neighbor.Y)
		if e == nil || !isHerbivore(e) {
			continue
		}

		prey, ok := e.(creature.Creature)
		if !ok {
			continue
		}

		return neighbor, prey, true
	}

	return coordinate.Point{}, nil, false
}

func isHerbivore(e entity.Entity) bool {
	_, ok := e.(*herbivore.Rabbit)
	return ok
}

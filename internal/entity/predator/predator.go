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

func (p *Predator) MakeMove(w *world.World) error {
	return p.BaseCreature.PerformMove(p, w)
}

func (p *Predator) HasAdjacentFood(w *world.World) (bool, error) {
	_, _, exists, err := p.findAdjacentFood(w)
	return exists, err
}

func (p *Predator) EatAdjacentFood(w *world.World) (bool, error) {
	foodPosition, prey, exists, err := p.findAdjacentFood(w)

	if err != nil {
		return false, err
	}

	if !exists {
		return false, nil
	}

	prey.TakeDamage(p.damage)
	if !prey.IsAlive() {
		if err = w.RemoveEntity(foodPosition); err != nil {
			return false, err
		}
	}

	return true, nil
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

func (p *Predator) findAdjacentFood(w *world.World) (coordinate.Point, creature.Creature, bool, error) {
	currentPosition, err := w.GetPointByEntity(p)
	if err != nil {
		return coordinate.Point{}, nil, false, err
	}

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

		return neighbor, prey, true, nil
	}

	return coordinate.Point{}, nil, false, nil
}

func isHerbivore(e entity.Entity) bool {
	_, ok := e.(*herbivore.Rabbit)
	return ok
}

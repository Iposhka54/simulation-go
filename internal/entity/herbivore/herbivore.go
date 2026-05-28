package herbivore

import (
	"simulation/internal/entity/creature"
	"simulation/internal/entity/static"
	"simulation/internal/game/path"
	"simulation/internal/game/world"
	"simulation/internal/game/world/coordinate"
)

type Herbivore struct {
	*creature.BaseCreature
}

func New(hp, maxHp, speed int) *Herbivore {
	return &Herbivore{
		BaseCreature: creature.New(hp, maxHp, speed),
	}
}

func (h *Herbivore) MakeMove(w *world.World) error {
	return h.BaseCreature.PerformMove(h, w)
}

func (h *Herbivore) HasAdjacentFood(w *world.World) (bool, error) {
	_, exists, err := h.findAdjacentFood(w)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (h *Herbivore) EatAdjacentFood(w *world.World) (bool, error) {
	foodPosition, exists, err := h.findAdjacentFood(w)
	if err != nil {
		return false, err
	}

	if !exists {
		return false, nil
	}

	if err = w.RemoveEntity(foodPosition); err != nil {
		return false, err
	}
	return true, nil
}

func (h *Herbivore) IsFoodAdjacent(w *world.World, p coordinate.Point) bool {
	for _, neighbor := range path.GetNeighbors(p) {
		if !w.IsValid(neighbor) {
			continue
		}

		e := w.Get(neighbor.X, neighbor.Y)
		if e == nil {
			continue
		}

		if _, ok := e.(*static.Grass); ok {
			return true
		}
	}

	return false
}

func (h *Herbivore) findAdjacentFood(w *world.World) (coordinate.Point, bool, error) {
	currentPosition, err := w.GetPointByEntity(h)
	if err != nil {
		return coordinate.Point{}, false, err
	}

	for _, neighbor := range path.GetNeighbors(currentPosition) {
		if !w.IsValid(neighbor) {
			continue
		}

		e := w.Get(neighbor.X, neighbor.Y)
		if e == nil {
			continue
		}

		if _, ok := e.(*static.Grass); ok {
			return neighbor, true, nil
		}
	}

	return coordinate.Point{}, false, nil
}

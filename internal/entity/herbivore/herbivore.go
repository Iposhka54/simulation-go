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

func (h *Herbivore) MakeMove(w *world.World) {
	h.BaseCreature.PerformMove(h, w)
}

func (h *Herbivore) HasAdjacentFood(w *world.World) bool {
	_, exists := h.findAdjacentFood(w)
	return exists
}

func (h *Herbivore) EatAdjacentFood(w *world.World) bool {
	foodPosition, exists := h.findAdjacentFood(w)
	if !exists {
		return false
	}

	w.RemoveEntity(foodPosition)
	return true
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

func (h *Herbivore) findAdjacentFood(w *world.World) (coordinate.Point, bool) {
	currentPosition := w.GetPointByEntity(h)
	for _, neighbor := range path.GetNeighbors(currentPosition) {
		if !w.IsValid(neighbor) {
			continue
		}

		e := w.Get(neighbor.X, neighbor.Y)
		if e == nil {
			continue
		}

		if _, ok := e.(*static.Grass); ok {
			return neighbor, true
		}
	}

	return coordinate.Point{}, false
}

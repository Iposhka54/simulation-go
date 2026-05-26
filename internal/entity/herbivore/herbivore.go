package herbivore

import (
	"simulation/internal/entity/creature"
	"simulation/internal/entity/static"
	_map "simulation/internal/game/map"
	"simulation/internal/game/map/coordinate"
	"simulation/internal/game/path"
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
	h.BaseCreature.PerformMove(h, m)
}

func (h *Herbivore) HasAdjacentFood(m *_map.Map) bool {
	_, exists := h.findAdjacentFood(m)
	return exists
}

func (h *Herbivore) EatAdjacentFood(m *_map.Map) bool {
	foodPosition, exists := h.findAdjacentFood(m)
	if !exists {
		return false
	}

	m.RemoveEntity(foodPosition)
	return true
}

func (h *Herbivore) IsFoodAdjacent(m *_map.Map, c coordinate.Point) bool {
	for _, neighbor := range path.GetNeighbors(c) {
		if !m.IsValid(neighbor) {
			continue
		}

		e := m.Get(neighbor.X, neighbor.Y)
		if e == nil {
			continue
		}

		if _, ok := e.(*static.Grass); ok {
			return true
		}
	}

	return false
}

func (h *Herbivore) findAdjacentFood(m *_map.Map) (coordinate.Point, bool) {
	currentPosition := m.GetCoordinatesByEntity(h)
	for _, neighbor := range path.GetNeighbors(currentPosition) {
		if !m.IsValid(neighbor) {
			continue
		}

		e := m.Get(neighbor.X, neighbor.Y)
		if e == nil {
			continue
		}

		if _, ok := e.(*static.Grass); ok {
			return neighbor, true
		}
	}

	return coordinate.Point{}, false
}

package action

import (
	"simulation/internal/entity/creature"
	_map "simulation/internal/game/map"
	"simulation/internal/game/map/coordinate"
)

type MoveAction struct{}

func (ma *MoveAction) Execute(worldMap *_map.Map) {
	positionedEntities := worldMap.GetPositionedEntities()

	for _, positioned := range positionedEntities {
		cr, ok := positioned.Entity.(creature.Creature)
		if !ok {
			continue
		}

		if !cr.IsAlive() {
			continue
		}

		if !isStillAtPosition(worldMap, positioned.Position, cr) {
			continue
		}

		cr.MakeMove(worldMap)
	}
}

func isStillAtPosition(worldMap *_map.Map, position coordinate.Coordinate, creature creature.Creature) bool {
	entity := worldMap.Get(position.X, position.Y)
	if entity == nil {
		return false
	}

	return entity == creature
}

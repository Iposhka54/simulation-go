package action

import (
	"simulation/internal/entity/creature"
	"simulation/internal/game/world"
	"simulation/internal/game/world/coordinate"
)

type MoveAction struct{}

func (ma *MoveAction) Execute(world *world.World) {
	positionedEntities := world.GetPositionedEntities()

	for _, positioned := range positionedEntities {
		cr, ok := positioned.Entity.(creature.Creature)
		if !ok {
			continue
		}

		if !cr.IsAlive() {
			continue
		}

		if !isStillAtPosition(world, positioned.Position, cr) {
			continue
		}

		cr.MakeMove(world)
	}
}

func isStillAtPosition(world *world.World, position coordinate.Point, creature creature.Creature) bool {
	entity := world.Get(position.X, position.Y)
	if entity == nil {
		return false
	}

	return entity == creature
}

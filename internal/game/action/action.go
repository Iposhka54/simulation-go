package action

import "simulation/internal/game/world"

type Action interface {
	Execute(world *world.World)
}

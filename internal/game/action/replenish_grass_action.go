package action

import (
	"math/rand"
	"simulation/internal/entity"
	"simulation/internal/entity/static"
	"simulation/internal/game/world"
)

type ReplenishGrassAction struct {
	probability     float64
	maxGrassToSpawn int
}

func NewReplenishGrassAction(probability float64, maxGrassToSpawn int) *ReplenishGrassAction {
	if probability < 0 {
		probability = 0
	}
	if probability > 1 {
		probability = 1
	}

	if maxGrassToSpawn < 1 {
		maxGrassToSpawn = 1
	}

	return &ReplenishGrassAction{
		probability:     probability,
		maxGrassToSpawn: maxGrassToSpawn,
	}
}

func (a *ReplenishGrassAction) Execute(world *world.World) {
	if rand.Float64() > a.probability {
		return
	}

	grassCount := rand.Intn(a.maxGrassToSpawn) + 1

	maxAttempts := world.Area() * MaxPlacementAttemptsMultiplier

	spawnUpTo(world, maxAttempts, grassCount, func() entity.Entity {
		return static.NewGrass()
	})
}

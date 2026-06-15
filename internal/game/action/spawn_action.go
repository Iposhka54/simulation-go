package action

import (
	"log"
	"math/rand"
	"simulation/internal/entity"
	"simulation/internal/entity/herbivore"
	"simulation/internal/entity/predator"
	"simulation/internal/entity/static"
	"simulation/internal/game/world"
	"simulation/internal/game/world/coordinate"
)

const (
	RockDensityDivisor             = 20
	TreeDensityDivisor             = 33
	GrassDensityDivisor            = 12
	HerbivoreDensityDivisor        = 25
	PredatorDensityDivisor         = 75
	MaxPlacementAttemptsMultiplier = 50

	InitialHerbivoreHp    = 50
	InitialHerbivoreSpeed = 1
	InitialPredatorHp     = 70
	PredatorSpeed         = 1
	PredatorDamage        = 8
)

type Spawner func() entity.Entity

type SpawnAction struct{}

func (sa *SpawnAction) Execute(world *world.World) {
	area := world.Area()
	maxAttempts := area * MaxPlacementAttemptsMultiplier

	spawnMany(world, area, PredatorDensityDivisor, maxAttempts, func() entity.Entity {
		return predator.NewWolf(InitialPredatorHp, InitialPredatorHp, PredatorSpeed, PredatorDamage)
	})
	spawnMany(world, area, HerbivoreDensityDivisor, maxAttempts, func() entity.Entity {
		return herbivore.NewRabbit(InitialHerbivoreHp, InitialHerbivoreHp, InitialHerbivoreSpeed)
	})
	spawnMany(world, area, RockDensityDivisor, maxAttempts, func() entity.Entity {
		return static.NewRock()
	})
	spawnMany(world, area, TreeDensityDivisor, maxAttempts, func() entity.Entity {
		return static.NewTree()
	})
	spawnMany(world, area, GrassDensityDivisor, maxAttempts, func() entity.Entity {
		return static.NewGrass()
	})
}

func spawnMany(world *world.World, area, divisor, maxAttempts int, spawner Spawner) {
	entityCount := area / divisor
	if entityCount < 1 {
		entityCount = 1
	}

	spawnUpTo(world, maxAttempts, entityCount, spawner)
}

func spawnUpTo(world *world.World, maxAttempts, entityCount int, spawner Spawner) {
	placed := 0
	attempts := 0
	for placed < entityCount && attempts < maxAttempts {
		spawnPosition := coordinate.New(rand.Intn(world.Width()), rand.Intn(world.Height()))
		if world.IsEmpty(spawnPosition) {
			if err := world.PlaceEntity(spawnPosition, spawner()); err != nil {
				log.Printf("spawn failed at %s: %v", spawnPosition.String(), err)
				continue
			}
			placed++
		}
		attempts++
	}
}

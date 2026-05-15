package action

import (
	"math/rand"
	"simulation/internal/entity"
	"simulation/internal/entity/herbivore"
	"simulation/internal/entity/predator"
	"simulation/internal/entity/static"
	_map "simulation/internal/game/map"
	"simulation/internal/game/map/coordinate"
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

func (sa *SpawnAction) Execute(worldMap *_map.Map) {
	area := worldMap.Area()
	maxAttempts := area * MaxPlacementAttemptsMultiplier

	spawnMany(worldMap, area, PredatorDensityDivisor, maxAttempts, func() entity.Entity {
		return predator.NewWolf(InitialPredatorHp, InitialPredatorHp, PredatorSpeed, PredatorDamage)
	})
	spawnMany(worldMap, area, HerbivoreDensityDivisor, maxAttempts, func() entity.Entity {
		return herbivore.NewRabbit(InitialHerbivoreHp, InitialHerbivoreHp, InitialHerbivoreSpeed)
	})
	spawnMany(worldMap, area, RockDensityDivisor, maxAttempts, func() entity.Entity {
		return static.NewRock()
	})
	spawnMany(worldMap, area, TreeDensityDivisor, maxAttempts, func() entity.Entity {
		return static.NewTree()
	})
	spawnMany(worldMap, area, GrassDensityDivisor, maxAttempts, func() entity.Entity {
		return static.NewGrass()
	})
}

func spawnMany(worldMap *_map.Map, area, divisor, maxAttempts int, spawner Spawner) {
	entityCount := area / divisor
	if entityCount < 1 {
		entityCount = 1
	}

	spawnUpTo(worldMap, maxAttempts, entityCount, spawner)
}

func spawnUpTo(worldMap *_map.Map, maxAttempts, entityCount int, spawner Spawner) {
	placed := 0
	attempts := 0
	for placed < entityCount && attempts < maxAttempts {
		spawnPosition := coordinate.New(rand.Intn(worldMap.Width()), rand.Intn(worldMap.Height()))
		if worldMap.IsEmpty(spawnPosition) {
			worldMap.PlaceEntity(spawnPosition, spawner())
			placed++
		}
		attempts++
	}
}

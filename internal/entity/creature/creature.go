package creature

import (
	"math/rand"
	"simulation/internal/entity"
	_map "simulation/internal/game/map"
	"simulation/internal/game/map/coordinate"
	"simulation/internal/game/path"
)

type Creature interface {
	entity.Entity
	MakeMove(m *_map.Map)
	Hp() int
	MaxHp() int
	Speed() int
	TakeDamage(damage int)
	IsAlive() bool
	Die(m *_map.Map)
	HasAdjacentFood(m *_map.Map) bool
	EatAdjacentFood(m *_map.Map) bool
	IsFoodAdjacent(m *_map.Map, c coordinate.Coordinate)
}

type BaseCreature struct {
	*entity.BaseEntity
	hp    int
	maxHp int
	speed int
}

func New(hp, maxHp, speed int) *BaseCreature {
	return &BaseCreature{
		hp:         hp,
		maxHp:      maxHp,
		speed:      speed,
		BaseEntity: entity.New(),
	}
}

func (bc *BaseCreature) MakeMove(m *_map.Map) {
	if !bc.IsAlive() {
		bc.Die(m)
	}

	bc.moveRandomly(m)
}

func (bc *BaseCreature) moveRandomly(m *_map.Map) {
	neighbors := path.FindReachableNeighbors(m, bc.getCurrentPosition(m))
	length := len(neighbors)
	if length >= 1 {
		i := rand.Intn(length)
		m.PlaceEntity(neighbors[i], bc)
		return
	}
	//need will log a situation where an entity cannot move
}

func (bc *BaseCreature) Hp() int {
	return bc.hp
}

func (bc *BaseCreature) MaxHp() int {
	return bc.maxHp
}

func (bc *BaseCreature) Speed() int {
	return bc.speed
}

func (bc *BaseCreature) TakeDamage(damage int) {
	bc.hp -= damage
	if bc.hp < 0 {
		bc.hp = 0
	}
}

func (bc *BaseCreature) IsAlive() bool {
	return bc.hp > 0
}

func (bc *BaseCreature) Die(m *_map.Map) {
	m.RemoveEntity(bc.getCurrentPosition(m))
}

func (bc *BaseCreature) HasAdjacentFood(m *_map.Map) bool {
	panic("implement in subclasses")
}

func (bc *BaseCreature) EatAdjacentFood(m *_map.Map) {
	panic("implement in subclasses")
}

func (bc *BaseCreature) IsFoodAdjacent(m *_map.Map, c coordinate.Coordinate) bool {
	panic("implement in subclasses")
}

func (bc *BaseCreature) getCurrentPosition(m *_map.Map) coordinate.Coordinate {
	return m.GetCoordinatesByEntity(bc)
}

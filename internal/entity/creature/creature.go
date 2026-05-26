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
	IsFoodAdjacent(m *_map.Map, c coordinate.Point) bool
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

func (bc *BaseCreature) PerformMove(c Creature, m *_map.Map) {
	if !bc.IsAlive() {
		bc.Die(m)
		return
	}

	if c.HasAdjacentFood(m) {
		if c.EatAdjacentFood(m) {
			return
		}
	}

	p := path.Find(m, bc.getCurrentPosition(m), c.IsFoodAdjacent)

	if len(p) > 1 {
		bc.moveAlongPath(m, p)
		return
	}

	bc.moveRandomly(m)
}

func (bc *BaseCreature) moveAlongPath(m *_map.Map, p []coordinate.Point) {
	length := len(p)
	if length <= 1 {
		panic("Path must contain at least 2 positions")
	}

	step := bc.speed
	if step >= length {
		step = length - 1
	}
	newPosition := p[step]
	bc.move(m, newPosition)
}

func (bc *BaseCreature) moveRandomly(m *_map.Map) {
	neighbors := path.FindReachableNeighbors(m, bc.getCurrentPosition(m))
	length := len(neighbors)
	if length >= 1 {
		i := rand.Intn(length)
		bc.move(m, neighbors[i])
		return
	}
	//todo need will log a situation where an entity cannot move
}

func (bc *BaseCreature) move(m *_map.Map, newPosition coordinate.Point) {
	position := bc.getCurrentPosition(m)
	movingEntity := m.Get(position.X, position.Y)
	m.RemoveEntity(position)
	m.PlaceEntity(newPosition, movingEntity)
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
	_ = m
	panic("implement in subclasses")
}

func (bc *BaseCreature) EatAdjacentFood(m *_map.Map) bool {
	_ = m
	panic("implement in subclasses")
}

func (bc *BaseCreature) IsFoodAdjacent(m *_map.Map, c coordinate.Point) bool {
	_, _ = m, c
	panic("implement in subclasses")
}

func (bc *BaseCreature) getCurrentPosition(m *_map.Map) coordinate.Point {
	return m.GetCoordinatesByEntity(bc)
}

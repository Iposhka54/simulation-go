package creature

import (
	"simulation/internal/entity"
	_map "simulation/internal/game/map"
)

type Creature interface {
	entity.Entity
	MakeMove(m *_map.Map)
	Hp() int
	MaxHp() int
	Speed() int
	TakeDamage(damage int)
	IsAlive() bool
	Eat(m *_map.Map)
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
	panic("implement in subclasses")
}

func (bc *BaseCreature) Eat(m *_map.Map) {
	panic("implement in subclasses")
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

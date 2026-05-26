package creature

import (
	"math/rand"
	"simulation/internal/entity"
	"simulation/internal/game/path"
	"simulation/internal/game/world"
	"simulation/internal/game/world/coordinate"
)

type Creature interface {
	entity.Entity
	MakeMove(w *world.World)
	Hp() int
	MaxHp() int
	Speed() int
	TakeDamage(damage int)
	IsAlive() bool
	Die(w *world.World)
	HasAdjacentFood(w *world.World) bool
	EatAdjacentFood(w *world.World) bool
	IsFoodAdjacent(w *world.World, c coordinate.Point) bool
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

func (bc *BaseCreature) PerformMove(c Creature, w *world.World) {
	if !bc.IsAlive() {
		bc.Die(w)
		return
	}

	if c.HasAdjacentFood(w) {
		if c.EatAdjacentFood(w) {
			return
		}
	}

	p := path.Find(w, bc.getCurrentPosition(w), c.IsFoodAdjacent)

	if len(p) > 1 {
		bc.moveAlongPath(w, p)
		return
	}

	bc.moveRandomly(w)
}

func (bc *BaseCreature) moveAlongPath(w *world.World, p []coordinate.Point) {
	length := len(p)
	if length <= 1 {
		panic("Path must contain at least 2 positions")
	}

	step := bc.speed
	if step >= length {
		step = length - 1
	}
	newPosition := p[step]
	bc.move(w, newPosition)
}

func (bc *BaseCreature) moveRandomly(w *world.World) {
	neighbors := path.FindReachableNeighbors(w, bc.getCurrentPosition(w))
	length := len(neighbors)
	if length >= 1 {
		i := rand.Intn(length)
		bc.move(w, neighbors[i])
		return
	}
	//todo need will log a situation where an entity cannot move
}

func (bc *BaseCreature) move(w *world.World, newPosition coordinate.Point) {
	position := bc.getCurrentPosition(w)
	movingEntity := w.Get(position.X, position.Y)
	w.RemoveEntity(position)
	w.PlaceEntity(newPosition, movingEntity)
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

func (bc *BaseCreature) Die(w *world.World) {
	w.RemoveEntity(bc.getCurrentPosition(w))
}

func (bc *BaseCreature) HasAdjacentFood(w *world.World) bool {
	_ = w
	panic("implement in subclasses")
}

func (bc *BaseCreature) EatAdjacentFood(w *world.World) bool {
	_ = w
	panic("implement in subclasses")
}

func (bc *BaseCreature) IsFoodAdjacent(w *world.World, p coordinate.Point) bool {
	_, _ = w, p
	panic("implement in subclasses")
}

func (bc *BaseCreature) getCurrentPosition(w *world.World) coordinate.Point {
	return w.GetPointByEntity(bc)
}

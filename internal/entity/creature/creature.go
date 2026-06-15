package creature

import (
	"fmt"
	"math/rand"
	"simulation/internal/entity"
	"simulation/internal/game/path"
	"simulation/internal/game/world"
	"simulation/internal/game/world/coordinate"
)

const HungerPerTurn = 3

type Creature interface {
	entity.Entity
	MakeMove(w *world.World) error
	Hp() int
	MaxHp() int
	Speed() int
	TakeDamage(damage int)
	IsAlive() bool
	Die(w *world.World) error
	HasAdjacentFood(w *world.World) (bool, error)
	EatAdjacentFood(w *world.World) (bool, error)
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

func (bc *BaseCreature) PerformMove(c Creature, w *world.World) error {
	bc.TakeDamage(HungerPerTurn)
	if !bc.IsAlive() {
		return bc.Die(w)
	}

	exists, err := c.HasAdjacentFood(w)
	if err != nil {
		return err
	}
	if exists {
		_, err = c.EatAdjacentFood(w)
		return err
	}

	position, err := bc.getCurrentPosition(w)
	if err != nil {
		return err
	}
	p := path.Find(w, position, c.IsFoodAdjacent)

	if len(p) > 1 {
		return bc.moveAlongPath(w, p)
	}

	if err = bc.moveRandomly(w); err != nil {
		return err
	}

	return nil
}

func (bc *BaseCreature) moveAlongPath(w *world.World, p []coordinate.Point) error {
	length := len(p)
	if length <= 1 {
		return fmt.Errorf("path must contain at least 2 positions, actual: %d", length)
	}

	step := bc.speed
	if step >= length {
		step = length - 1
	}
	newPosition := p[step]
	return bc.move(w, newPosition)
}

func (bc *BaseCreature) moveRandomly(w *world.World) error {
	position, err := bc.getCurrentPosition(w)
	if err != nil {
		return err
	}
	neighbors := path.FindReachableNeighbors(w, position)
	length := len(neighbors)
	if length >= 1 {
		i := rand.Intn(length)
		return bc.move(w, neighbors[i])
	}
	//todo need will log a situation where an entity cannot move

	return nil
}

func (bc *BaseCreature) move(w *world.World, newPosition coordinate.Point) error {
	position, err := bc.getCurrentPosition(w)
	if err != nil {
		return err
	}

	movingEntity := w.Get(position.X, position.Y)
	if err = w.RemoveEntity(position); err != nil {
		return err
	}

	if err = w.PlaceEntity(newPosition, movingEntity); err != nil {
		return err
	}
	return nil
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

func (bc *BaseCreature) Die(w *world.World) error {
	position, err := bc.getCurrentPosition(w)
	if err != nil {
		return err
	}

	return w.RemoveEntity(position)
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

func (bc *BaseCreature) getCurrentPosition(w *world.World) (coordinate.Point, error) {
	return w.GetPointByEntity(bc)
}

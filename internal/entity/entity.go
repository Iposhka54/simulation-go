package entity

import (
	"simulation/internal/game/map/coordinate"
)

type Entity interface {
	MakeMove(to coordinate.Coordinate)
	SetCoordinates(c coordinate.Coordinate)
	Coordinates() coordinate.Coordinate
	String() string
	Hp() int
	SetHp(hp int)
}

type BaseEntity struct {
	coordinate coordinate.Coordinate
}

func New(c coordinate.Coordinate) *BaseEntity {
	return &BaseEntity{coordinate: c}
}

func (e *BaseEntity) MakeMove(to coordinate.Coordinate) {
	panic("need override in the heirs")
}

func (e *BaseEntity) SetCoordinates(to coordinate.Coordinate) {
	e.coordinate = to
}

func (e *BaseEntity) Coordinates() coordinate.Coordinate {
	return e.coordinate
}

func (e *BaseEntity) String() string {
	panic("need override in the heirs")
}

func (e *BaseEntity) Hp() int {
	return 0
}

func (e *BaseEntity) SetHp(hp int) {

}

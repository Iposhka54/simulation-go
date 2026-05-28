package world

import (
	"errors"
	"fmt"
	"simulation/internal/entity"
	"simulation/internal/game/world/coordinate"
)

var (
	ErrOutOfBounds     = errors.New("point are out of bounds")
	ErrAlreadyOccupied = errors.New("cell is already occupied")
	ErrCellEmpty       = errors.New("cannot remove entity because cell is empty")
	ErrEntityNotFound  = errors.New("entity not found on the map")
)

type World struct {
	width              int
	height             int
	entitiesByPoints   map[coordinate.Point]entity.Entity
	pointsByEntitiesID map[uint64]coordinate.Point
}

type PositionedEntity struct {
	Entity   entity.Entity
	Position coordinate.Point
}

func New(width, height int) (*World, error) {
	if err := validateWorldSize(width, height); err != nil {
		return nil, err
	}

	return &World{
		width:              width,
		height:             height,
		entitiesByPoints:   make(map[coordinate.Point]entity.Entity),
		pointsByEntitiesID: make(map[uint64]coordinate.Point),
	}, nil
}

func (w *World) PlaceEntity(p coordinate.Point, e entity.Entity) error {
	if !w.IsValid(p) {
		return fmt.Errorf("%w: %s for world %dx%d", ErrOutOfBounds, p.String(), w.width, w.height)
	}
	if !w.IsEmpty(p) {
		return fmt.Errorf("%w: %s", ErrAlreadyOccupied, p.String())
	}
	w.entitiesByPoints[p] = e
	w.pointsByEntitiesID[e.ID()] = p

	return nil
}

func (w *World) IsEmpty(c coordinate.Point) bool {
	if _, exists := w.entitiesByPoints[c]; exists {
		return false
	}
	return true
}

func (w *World) Width() int {
	return w.width
}

func (w *World) Height() int {
	return w.height
}

func (w *World) Area() int {
	return w.height * w.width
}

func (w *World) Get(x, y int) entity.Entity {
	c := coordinate.New(x, y)
	e, ok := w.entitiesByPoints[c]
	if !ok {
		return nil
	}
	return e
}

func (w *World) GetPointByEntity(e entity.Entity) (coordinate.Point, error) {
	point, exists := w.pointsByEntitiesID[e.ID()]
	if !exists {
		return coordinate.Point{}, fmt.Errorf("%w: ID %d", ErrEntityNotFound, e.ID())
	}
	return point, nil
}

func (w *World) IsValid(c coordinate.Point) bool {
	return c.X >= 0 && c.X < w.width && c.Y >= 0 && c.Y < w.height
}

func (w *World) RemoveEntity(p coordinate.Point) error {
	if !w.IsValid(p) {
		return fmt.Errorf("%w: %s for world %dx%d", ErrOutOfBounds, p.String(), w.width, w.height)
	}
	if w.IsEmpty(p) {
		return fmt.Errorf("%w: %s", ErrCellEmpty, p.String())
	}

	e := w.entitiesByPoints[p]

	delete(w.entitiesByPoints, p)
	delete(w.pointsByEntitiesID, e.ID())

	return nil
}

func (w *World) GetPositionedEntities() []PositionedEntity {
	positioned := make([]PositionedEntity, 0, len(w.entitiesByPoints))

	for point, e := range w.entitiesByPoints {
		positioned = append(positioned, PositionedEntity{
			Entity:   e,
			Position: point,
		})
	}

	return positioned
}

func validateWorldSize(width, height int) error {
	if width <= 0 || height <= 0 {
		return fmt.Errorf("invalid world size: width=%d, height=%d", width, height)
	}

	return nil
}

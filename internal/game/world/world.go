package world

import (
	"fmt"
	"simulation/internal/entity"
	"simulation/internal/game/world/coordinate"
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

func New(width, height int) World {
	validateWorldSize(width, height)
	return World{
		width:              width,
		height:             height,
		entitiesByPoints:   make(map[coordinate.Point]entity.Entity),
		pointsByEntitiesID: make(map[uint64]coordinate.Point),
	}
}

func (w *World) PlaceEntity(p coordinate.Point, e entity.Entity) {
	w.validate(p)
	if !w.IsEmpty(p) {
		panic(fmt.Sprintf("cell %s is already occupied", p.String()))
	}
	w.entitiesByPoints[p] = e
	w.pointsByEntitiesID[e.ID()] = p
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

func (w *World) GetPointByEntity(e entity.Entity) coordinate.Point {
	point, exists := w.pointsByEntitiesID[e.ID()]
	if !exists {
		panic(fmt.Sprintf("Entity %v not found on the world", e))
	}
	return point
}

func (w *World) IsValid(c coordinate.Point) bool {
	return c.X >= 0 && c.X < w.width && c.Y >= 0 && c.Y < w.height
}

func (w *World) RemoveEntity(p coordinate.Point) {
	w.validate(p)

	if w.IsEmpty(p) {
		panic(fmt.Sprintf("Cannot remove entity from %s because this point is empty", p.String()))
	}

	e := w.entitiesByPoints[p]

	delete(w.entitiesByPoints, p)
	delete(w.pointsByEntitiesID, e.ID())
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

func (w *World) validate(p coordinate.Point) {
	if w.IsValid(p) {
		return
	}
	panic(fmt.Sprintf("%s are out of bounds for world %dx%d", p.String(), w.width, w.height))
}

func validateWorldSize(width, height int) {
	if width <= 0 || height <= 0 {
		panic(fmt.Sprintf("Invalid world size: width=%d, height=%d", width, height))
	}
}

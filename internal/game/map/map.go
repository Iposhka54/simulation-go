package _map

import (
	"fmt"
	"simulation/internal/entity"
	"simulation/internal/game/map/coordinate"
)

type Map struct {
	width                 int
	height                int
	entitiesByCoordinates map[coordinate.Point]entity.Entity
	coordinateByEntityID  map[uint64]coordinate.Point
}

type PositionedEntity struct {
	Entity   entity.Entity
	Position coordinate.Point
}

func New(width, height int) Map {
	validateMapSize(width, height)
	return Map{
		width:                 width,
		height:                height,
		entitiesByCoordinates: make(map[coordinate.Point]entity.Entity),
		coordinateByEntityID:  make(map[uint64]coordinate.Point),
	}
}

func (m *Map) PlaceEntity(c coordinate.Point, e entity.Entity) {
	m.validate(c)
	if !m.IsEmpty(c) {
		panic(fmt.Sprintf("cell %s is already occupied", c.String()))
	}
	m.entitiesByCoordinates[c] = e
	m.coordinateByEntityID[e.ID()] = c
}

func (m *Map) IsEmpty(c coordinate.Point) bool {
	if _, exists := m.entitiesByCoordinates[c]; exists {
		return false
	}
	return true
}

func (m *Map) Width() int {
	return m.width
}

func (m *Map) Height() int {
	return m.height
}

func (m *Map) Area() int {
	return m.height * m.width
}

func (m *Map) Get(x, y int) entity.Entity {
	c := coordinate.New(x, y)
	e, ok := m.entitiesByCoordinates[c]
	if !ok {
		return nil
	}
	return e
}

func (m *Map) GetCoordinatesByEntity(e entity.Entity) coordinate.Point {
	coord, exists := m.coordinateByEntityID[e.ID()]
	if !exists {
		panic(fmt.Sprintf("Entity %v not found on the map", e))
	}
	return coord
}

func (m *Map) IsValid(c coordinate.Point) bool {
	return c.X >= 0 && c.X < m.width && c.Y >= 0 && c.Y < m.height
}

func (m *Map) RemoveEntity(c coordinate.Point) {
	m.validate(c)

	if m.IsEmpty(c) {
		panic(fmt.Sprintf("Cannot remove entity from %s because this coords is empty", c.String()))
	}

	e := m.entitiesByCoordinates[c]

	delete(m.entitiesByCoordinates, c)
	delete(m.coordinateByEntityID, e.ID())
}

func (m *Map) GetPositionedEntities() []PositionedEntity {
	positioned := make([]PositionedEntity, 0, len(m.entitiesByCoordinates))

	for coord, e := range m.entitiesByCoordinates {
		positioned = append(positioned, PositionedEntity{
			Entity:   e,
			Position: coord,
		})
	}

	return positioned
}

func (m *Map) validate(c coordinate.Point) {
	if m.IsValid(c) {
		return
	}
	panic(fmt.Sprintf("Coordinates %s are out of bounds for map %dx%d", c.String(), m.width, m.height))
}

func validateMapSize(width, height int) {
	if width <= 0 || height <= 0 {
		panic(fmt.Sprintf("Invalid map size: width=%d, height=%d", width, height))
	}
}

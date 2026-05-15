package _map

import (
	"fmt"
	"simulation/internal/entity"
	"simulation/internal/game/map/coordinate"
)

type Map struct {
	width                 int
	height                int
	entitiesByCoordinates map[coordinate.Coordinate]entity.Entity
	coordinateByEntities  map[entity.Entity]coordinate.Coordinate
}

func New(width, height int) Map {
	validateMapSize(width, height)
	return Map{
		width:                 width,
		height:                height,
		entitiesByCoordinates: make(map[coordinate.Coordinate]entity.Entity),
		coordinateByEntities:  make(map[entity.Entity]coordinate.Coordinate),
	}
}

func (m *Map) PlaceEntity(c coordinate.Coordinate, e entity.Entity) {
	m.validate(c)
	if !m.IsEmpty(c) {
		panic(fmt.Sprintf("cell %s is already occupied", c.String()))
	}
	m.entitiesByCoordinates[c] = e
	m.coordinateByEntities[e] = c
}

func (m *Map) IsEmpty(c coordinate.Coordinate) bool {
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

func (m *Map) IsValid(c coordinate.Coordinate) bool {
	return c.X > 0 && c.X < m.width && c.Y > 0 && c.Y < m.height
}

func (m *Map) validate(c coordinate.Coordinate) {
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

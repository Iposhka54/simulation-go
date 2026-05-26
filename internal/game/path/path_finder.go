package path

import (
	_map "simulation/internal/game/map"
	"simulation/internal/game/map/coordinate"
)

var directions = []coordinate.Point{
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
}

type HasAdjacentFood func(worldMap *_map.Map, c coordinate.Point) bool

func Find(worldMap *_map.Map,
	startPosition coordinate.Point,
	foodChecker HasAdjacentFood) []coordinate.Point {
	if foodChecker(worldMap, startPosition) {
		return []coordinate.Point{startPosition}
	}

	queue := []coordinate.Point{startPosition}
	visited := make(map[coordinate.Point]bool)
	parents := make(map[coordinate.Point]coordinate.Point)

	visited[startPosition] = true

	for len(queue) > 0 {
		currentPosition := queue[0]
		queue = queue[1:]

		neighbors := FindReachableNeighbors(worldMap, currentPosition)

		for _, neighbor := range neighbors {
			if visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			parents[neighbor] = currentPosition

			if foodChecker(worldMap, neighbor) {
				return reconstructPath(neighbor, startPosition, parents)
			}

			queue = append(queue, neighbor)
		}
	}

	return []coordinate.Point{}
}

func FindReachableNeighbors(worldMap *_map.Map,
	position coordinate.Point) []coordinate.Point {
	neighbors := GetNeighbors(position)

	filtered := neighbors[:0]

	for _, neighbor := range neighbors {
		if worldMap.IsValid(neighbor) && worldMap.IsEmpty(neighbor) {
			filtered = append(filtered, neighbor)
		}
	}

	return filtered
}

func GetNeighbors(position coordinate.Point) []coordinate.Point {
	var neighbors []coordinate.Point

	for _, dir := range directions {
		neighbors = append(neighbors, coordinate.Point{
			X: position.X + dir.X,
			Y: position.Y + dir.Y,
		})
	}

	return neighbors
}

func reconstructPath(goalCord, startCord coordinate.Point,
	parents map[coordinate.Point]coordinate.Point) []coordinate.Point {
	path := []coordinate.Point{goalCord}

	currentCord := goalCord
	for currentCord != startCord {
		currentCord = parents[currentCord]
		path = append([]coordinate.Point{currentCord}, path...)
	}

	return path
}

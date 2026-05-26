package path

import (
	_map "simulation/internal/game/map"
	"simulation/internal/game/map/coordinate"
)

var directions = []coordinate.Coordinate{
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
}

type HasAdjacentFood func(worldMap *_map.Map, c coordinate.Coordinate) bool

func Find(worldMap *_map.Map,
	startPosition coordinate.Coordinate,
	foodChecker HasAdjacentFood) []coordinate.Coordinate {
	if foodChecker(worldMap, startPosition) {
		return []coordinate.Coordinate{startPosition}
	}

	queue := []coordinate.Coordinate{startPosition}
	visited := make(map[coordinate.Coordinate]bool)
	parents := make(map[coordinate.Coordinate]coordinate.Coordinate)

	visited[startPosition] = true

	for len(queue) > 0 {
		currentPosition := queue[0]
		queue = queue[1:]

		neighbors := FindReachableNeighbors(worldMap, currentPosition)

		for _, neighbor := range neighbors {
			if visited[neighbor] {
				continue
			}

			parents[neighbor] = currentPosition

			if foodChecker(worldMap, neighbor) {
				return reconstructPath(neighbor, startPosition, parents)
			}

			queue = append(queue, neighbor)
		}
	}

	return []coordinate.Coordinate{}
}

func FindReachableNeighbors(worldMap *_map.Map,
	position coordinate.Coordinate) []coordinate.Coordinate {
	neighbors := GetNeighbors(position)

	filtered := neighbors[:0]

	for _, neighbor := range neighbors {
		if worldMap.IsValid(neighbor) && worldMap.IsEmpty(neighbor) {
			filtered = append(filtered, neighbor)
		}
	}

	return filtered
}

func GetNeighbors(position coordinate.Coordinate) []coordinate.Coordinate {
	var neighbors []coordinate.Coordinate

	for _, dir := range directions {
		neighbors = append(neighbors, coordinate.Coordinate{
			X: position.X + dir.X,
			Y: position.Y + dir.Y,
		})
	}

	return neighbors
}

func reconstructPath(goalCord, startCord coordinate.Coordinate,
	parents map[coordinate.Coordinate]coordinate.Coordinate) []coordinate.Coordinate {
	path := []coordinate.Coordinate{goalCord}

	currentCord := goalCord
	for currentCord != startCord {
		currentCord = parents[currentCord]
		path = append([]coordinate.Coordinate{currentCord}, path...)
	}

	return path
}

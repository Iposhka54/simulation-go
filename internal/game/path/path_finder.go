package path

import (
	"simulation/internal/game/world"
	"simulation/internal/game/world/coordinate"
)

var directions = []coordinate.Point{
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
}

type HasAdjacentFood func(world *world.World, p coordinate.Point) bool

func Find(world *world.World,
	startPosition coordinate.Point,
	foodChecker HasAdjacentFood) []coordinate.Point {
	if foodChecker(world, startPosition) {
		return []coordinate.Point{startPosition}
	}

	queue := []coordinate.Point{startPosition}
	visited := make(map[coordinate.Point]bool)
	parents := make(map[coordinate.Point]coordinate.Point)

	visited[startPosition] = true

	for len(queue) > 0 {
		currentPosition := queue[0]
		queue = queue[1:]

		neighbors := FindReachableNeighbors(world, currentPosition)

		for _, neighbor := range neighbors {
			if visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			parents[neighbor] = currentPosition

			if foodChecker(world, neighbor) {
				return reconstructPath(neighbor, startPosition, parents)
			}

			queue = append(queue, neighbor)
		}
	}

	return []coordinate.Point{}
}

func FindReachableNeighbors(world *world.World,
	position coordinate.Point) []coordinate.Point {
	neighbors := GetNeighbors(position)

	filtered := neighbors[:0]

	for _, neighbor := range neighbors {
		if world.IsValid(neighbor) && world.IsEmpty(neighbor) {
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

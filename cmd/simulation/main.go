package main

import (
	"simulation/internal/game/action"
	_map "simulation/internal/game/map"
	"simulation/internal/game/renderer"
	"simulation/internal/game/renderer/glyph_set"
	"simulation/internal/game/simulation"
)

const (
	DefaultWidth   = 10
	DefaultHeight  = 10
	DefaultDelayMs = 1000
)

func main() {
	worldMap := _map.New(DefaultWidth, DefaultHeight)

	initActions := []action.Action{
		&action.SpawnAction{},
	}

	turnActions := []action.Action{}

	r := renderer.NewConsoleRenderer(renderer.EmptyCellGlyph, glyph_set.NewEmojiGlyphSet())
	s := simulation.New(&worldMap, r, initActions, turnActions)

	s.PrintSimulation()
}

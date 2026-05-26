package main

import (
	"simulation/internal/app"
	"simulation/internal/game/action"
	"simulation/internal/game/renderer"
	"simulation/internal/game/renderer/glyph_set"
	"simulation/internal/game/simulation"
	"simulation/internal/game/world"
)

const (
	DefaultWidth   = 10
	DefaultHeight  = 10
	DefaultDelayMs = 1000
)

func main() {
	w := world.New(DefaultWidth, DefaultHeight)

	initActions := []action.Action{
		&action.SpawnAction{},
	}

	turnActions := []action.Action{
		&action.MoveAction{},
	}

	r := renderer.NewConsoleRenderer(renderer.EmptyCellGlyph, glyph_set.NewEmojiGlyphSet())
	s := simulation.New(&w, DefaultDelayMs, r, initActions, turnActions)
	a := app.New(s)
	a.Run()
}

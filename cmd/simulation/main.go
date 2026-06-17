package main

import (
	"context"
	"log"
	"os"
	"os/signal"
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
	ctx := context.Background()
	ctx, cancelFunc := signal.NotifyContext(ctx, os.Interrupt)
	w, err := world.New(DefaultWidth, DefaultHeight)
	if err != nil {
		log.Fatalf("Critical initialization world error: %v", err)
	}

	initActions := []action.Action{
		&action.SpawnAction{},
	}

	turnActions := []action.Action{
		&action.MoveAction{},
		action.NewReplenishGrassAction(0.2, 3),
	}

	r := renderer.NewConsoleRenderer(renderer.EmptyCellGlyph, glyph_set.NewEmojiGlyphSet())
	s := simulation.New(w, DefaultDelayMs, r, initActions, turnActions)
	a := app.New(cancelFunc, s)
	a.Run(ctx)
}

package renderer

import "simulation/internal/game/world"

const (
	EmptyCellGlyph = "⬛"
)

type Renderer interface {
	Render(world *world.World)
}

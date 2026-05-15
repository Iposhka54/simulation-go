package renderer

import _map "simulation/internal/game/map"

const (
	EmptyCellGlyph = "⬛"
)

type Renderer interface {
	Render(worldMap *_map.Map)
}

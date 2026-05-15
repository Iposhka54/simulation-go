package renderer

import (
	_map "simulation/internal/game/map"
	"simulation/internal/game/renderer/glyph_set"
	"strings"
)

type ConsoleRenderer struct {
	emptyCellGlyph string
	glyphSet       glyph_set.GlyphSet
}

func NewConsoleRenderer(emptyCellGlyph string, glyphSet glyph_set.GlyphSet) *ConsoleRenderer {
	return &ConsoleRenderer{
		emptyCellGlyph: emptyCellGlyph,
		glyphSet:       glyphSet,
	}
}

func (cr *ConsoleRenderer) Render(worldMap *_map.Map) {
	output := strings.Builder{}

	for y := 0; y < worldMap.Height(); y++ {
		for x := 0; x < worldMap.Width(); x++ {
			entity := worldMap.Get(x, y)
			if entity != nil {
				output.WriteString(cr.glyphSet.GetGlyph(entity))
			} else {
				output.WriteString(cr.emptyCellGlyph)
			}
			output.WriteString(" ")
		}
		output.WriteString("\n")
	}
	print(output.String())
}

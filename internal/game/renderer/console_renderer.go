package renderer

import (
	"simulation/internal/game/renderer/glyph_set"
	"simulation/internal/game/world"
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

func (cr *ConsoleRenderer) Render(world *world.World) {
	output := strings.Builder{}

	for y := 0; y < world.Height(); y++ {
		for x := 0; x < world.Width(); x++ {
			entity := world.Get(x, y)
			if entity != nil {
				output.WriteString(cr.glyphSet.GetGlyph(entity))
			} else {
				output.WriteString(cr.emptyCellGlyph)
			}
			output.WriteString(" ")
		}
		output.WriteString("\n")
	}
	println(output.String())
}

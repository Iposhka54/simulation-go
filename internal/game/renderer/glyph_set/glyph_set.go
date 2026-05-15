package glyph_set

import "simulation/internal/entity"

type GlyphSet interface {
	GetGlyph(entity entity.Entity) string
}

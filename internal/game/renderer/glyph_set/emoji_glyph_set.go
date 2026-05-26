package glyph_set

import (
	"simulation/internal/entity"
	"simulation/internal/entity/herbivore"
	"simulation/internal/entity/predator"
	"simulation/internal/entity/static"
)

type EmojiGlyphSet struct{}

func NewEmojiGlyphSet() *EmojiGlyphSet {
	return &EmojiGlyphSet{}
}

func (e *EmojiGlyphSet) GetGlyph(entity entity.Entity) string {
	switch entity.(type) {
	case *static.Grass:
		return "🌿"
	case *static.Tree:
		return "🌳"
	case *static.Rock:
		return "🪨"
	case *herbivore.Rabbit:
		return "🐇"
	case *predator.Wolf:
		return "🐺"
	default:
		return "?"
	}
}

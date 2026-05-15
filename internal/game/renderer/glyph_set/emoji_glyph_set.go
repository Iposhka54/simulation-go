package glyph_set

import (
	"reflect"
	"simulation/internal/entity"
	"simulation/internal/entity/herbivore"
	"simulation/internal/entity/predator"
	"simulation/internal/entity/static"
)

type EmojiGlyphSet struct {
	glyphs map[reflect.Type]string
}

func NewEmojiGlyphSet() *EmojiGlyphSet {
	return &EmojiGlyphSet{
		glyphs: map[reflect.Type]string{
			reflect.PointerTo(reflect.TypeOf(static.Grass{})):     "🌿",
			reflect.PointerTo(reflect.TypeOf(static.Tree{})):      "🌳",
			reflect.PointerTo(reflect.TypeOf(static.Rock{})):      "🪨",
			reflect.PointerTo(reflect.TypeOf(herbivore.Rabbit{})): "🐇",
			reflect.PointerTo(reflect.TypeOf(predator.Wolf{})):    "🐺",
		},
	}
}

func (e *EmojiGlyphSet) GetGlyph(entity entity.Entity) string {
	glyph, ok := e.glyphs[reflect.TypeOf(entity)]
	if !ok {
		return "?"
	}
	return glyph
}

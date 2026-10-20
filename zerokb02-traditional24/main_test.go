package main

import (
	"testing"

	"github.com/rdon-key/modgadget-fonts/efont24"
)

func TestFullEfont24GreetingGlyphCoverage(t *testing.T) {
	font := efont24.Font
	if !font.Valid() {
		t.Fatal("full Efont24 font is invalid")
	}
	for _, r := range "恭喜進軍台灣！" {
		if !font.HasGlyph(r) {
			t.Errorf("full Efont24 is missing glyph %q (U+%04X)", r, r)
		}
	}
}

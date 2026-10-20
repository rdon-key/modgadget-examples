package main

import (
	"testing"

	"github.com/rdon-key/modgadget-fonts/efont16"
)

// Test the full Efont16 font, not a greeting-only subset, on the Go host.
func TestGreetingGlyphCoverage(t *testing.T) {
	const greeting = "恭喜進軍台灣！"
	font := efont16.Font
	if !font.Valid() {
		t.Fatal("full Efont16 font is invalid")
	}
	for _, r := range greeting {
		if !font.HasGlyph(r) {
			t.Errorf("full Efont16 is missing glyph %q (U+%04X)", r, r)
		}
	}
}

package main

import (
	"testing"

	"github.com/rdon-key/modgadget-fonts/efont16zerokb02"
)

// Catch a missing character on the host before a TinyGo build produces
// a silent panic and a blank OLED.
func TestGreetingGlyphCoverage(t *testing.T) {
	const greeting = "恭喜進軍台灣！"
	if !efont16zerokb02.Font.Valid() {
		t.Fatal("efont16 font is invalid")
	}
	for _, r := range greeting {
		if !efont16zerokb02.Font.HasGlyph(r) {
			t.Errorf("efont16zerokb02 is missing glyph %q (U+%04X)", r, r)
		}
	}
}

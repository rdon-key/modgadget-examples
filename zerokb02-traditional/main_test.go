package main

import (
	"testing"

	"github.com/rdon-key/modgadget-fonts/efont16"
)

// Catch a missing character on the host before a TinyGo build produces
// a silent panic and a blank OLED.
func TestGreetingGlyphCoverage(t *testing.T) {
	const greeting = "恭喜進軍台灣！"
	if !efont16.Font.Valid() {
		t.Fatal("efont16 font is invalid")
	}
	for _, r := range greeting {
		if !efont16.Font.HasGlyph(r) {
			t.Errorf("efont16 is missing glyph %q (U+%04X)", r, r)
		}
	}
}

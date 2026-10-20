//go:build tinygo

package main

import (
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/efont16"
	board "github.com/rdon-key/modgadget/device/zerokb02"
)

// This demo uses an external compressed MGZ font; the Zero-KB02-specific
// hardware code remains in ModGadget's device/zerokb02 package.
func main() {
	// Leave time for the USB CDC serial monitor to attach.
	time.Sleep(2 * time.Second)
	println("zero-kb02: greeting demo starting")
	panel, err := board.ConfigureDisplay()
	if err != nil {
		halt("OLED init", err.Error())
	}
	println("zero-kb02: OLED configured")

	font := efont16.Font
	if !font.Valid() {
		halt("invalid font", "efont16 font handle is invalid")
	}
	println("zero-kb02: Efont16 initialized")
	const message = "恭喜進軍台灣！"
	for _, r := range message {
		println("zero-kb02: lookup glyph:", string(r))
		if !font.HasGlyph(r) {
			// HasGlyph calls the full MGZ lookup/decompression path.
			halt("glyph lookup", string(r))
		}
	}
	println("zero-kb02: all glyphs decoded")

	gadget := modgadget.New(panel, modgadget.WithStyles(modgadget.StyleSet{
		Default: modgadget.Style{
			Font:       font,
			Foreground: modgadget.ColorWhite,
			Background: modgadget.ColorBlack,
		},
	}))
	println("zero-kb02: clearing OLED")
	if err := gadget.Clear(); err != nil {
		halt("OLED clear", err.Error())
	}
	println("zero-kb02: OLED cleared")

	// Seven 16-dot CJK glyphs fit across the 128-pixel OLED.
	// Keep the complete greeting on one line, approximately centered.
	view := gadget.Viewport(modgadget.Bounds(8, 20, 120, 24))
	println("zero-kb02: building text layout")
	if err := view.SetText(message); err != nil {
		halt("SetText", err.Error())
	}
	println("zero-kb02: rendering text")
	if err := gadget.Render(); err != nil {
		halt("Render", err.Error())
	}
	println("zero-kb02: render complete")

	for {
		println("zero-kb02: heartbeat, render complete")
		time.Sleep(2 * time.Second)
	}
}

// Repeat the failure even if the USB serial monitor attaches after startup.
func halt(stage, detail string) {
	for {
		println("zero-kb02: FAILED at", stage, ":", detail)
		time.Sleep(2 * time.Second)
	}
}

//go:build tinygo

package main

import (
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/efont24"
	board "github.com/rdon-key/modgadget/device/zerokb02"
)

// This demo uses an external compressed MGZ font; the Zero-KB02-specific
// hardware code remains in ModGadget's device/zerokb02 package.
func main() {
	panel, err := board.ConfigureDisplay()
	if err != nil {
		panic(err)
	}

	font := efont24.Font
	const message = "恭喜進軍台灣！"
	for _, r := range message {
		if !font.HasGlyph(r) {
			panic("efont24: missing greeting glyph")
		}
	}

	gadget := modgadget.New(panel, modgadget.WithStyles(modgadget.StyleSet{
		Default: modgadget.Style{
			Font:       font,
			Foreground: modgadget.ColorWhite,
			Background: modgadget.ColorBlack,
		},
	}))
	if err := gadget.Clear(); err != nil {
		panic(err)
	}

	// Four 24-dot glyphs fit on the first line of the 128-pixel OLED.
	first := gadget.Viewport(modgadget.Bounds(16, 0, 112, 32))
	if err := first.SetText("恭喜進軍"); err != nil {
		panic(err)
	}
	second := gadget.Viewport(modgadget.Bounds(28, 32, 100, 32))
	if err := second.SetText("台灣！"); err != nil {
		panic(err)
	}
	if err := gadget.Render(); err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second)
	}
}

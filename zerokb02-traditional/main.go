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
	panel, err := board.ConfigureDisplay()
	if err != nil {
		panic(err)
	}

	font := efont16.Font
	const message = "恭喜進軍台灣！"
	for _, r := range message {
		if !font.HasGlyph(r) {
			panic("efont16: missing greeting glyph")
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

	// Seven 16-dot CJK glyphs fit across the 128-pixel OLED.
	// Keep the complete greeting on one line, approximately centered.
	view := gadget.Viewport(modgadget.Bounds(8, 20, 120, 24))
	if err := view.SetText(message); err != nil {
		panic(err)
	}
	if err := gadget.Render(); err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second)
	}
}

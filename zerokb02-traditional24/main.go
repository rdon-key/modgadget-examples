//go:build tinygo

package main

import (
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/efont24"
	board "github.com/rdon-key/modgadget/device/zerokb02"
)

// Efont Biwidth 24 contains the entire distributed MGZ font, not a subset.
// This program requires the 2 MiB Zero-KB02 target in ../targets/.
const greeting = "恭喜進軍台灣！"

func main() {
	// The font package is initialized before main. USB CDC monitoring begins
	// here after successful package initialization.
	time.Sleep(2 * time.Second)
	println("zero-kb02 24: starting full Efont24 horizontal scroll")

	panel, err := board.ConfigureDisplay()
	if err != nil {
		halt("OLED init", err.Error())
	}
	println("zero-kb02 24: OLED configured")

	font := efont24.Font
	if !font.Valid() {
		halt("font", "full Efont24 handle is invalid")
	}
	println("zero-kb02 24: Efont24 initialized")

	for _, r := range greeting {
		println("zero-kb02 24: lookup glyph:", string(r))
		if !font.HasGlyph(r) {
			halt("glyph lookup", string(r))
		}
	}
	println("zero-kb02 24: all seven glyphs decoded")

	gadget := modgadget.New(panel, modgadget.WithStyles(modgadget.StyleSet{
		Default: modgadget.Style{
			Font:       font,
			Foreground: modgadget.ColorWhite,
			Background: modgadget.ColorBlack,
		},
	}))
	if err := gadget.Clear(); err != nil {
		halt("OLED clear", err.Error())
	}

	// 24-dot Chinese characters are wider than a 128-pixel display.
	// A single viewport clips the line; horizontal scrolling reveals the rest.
	view := gadget.Viewport(modgadget.Bounds(0, 16, 128, 32))
	if err := view.SetText(greeting); err != nil {
		halt("SetText", err.Error())
	}
	view.SetHorizontalScroll(
		modgadget.ScrollSpeed(24),
		modgadget.ScrollGap(32),
		modgadget.ScrollLoop(),
		modgadget.ScrollFromRight(),
	)
	println("zero-kb02 24: starting horizontal marquee")

	nextHeartbeat := time.Now().Add(2 * time.Second)
	for {
		gadget.Update(time.Now())
		if err := gadget.Render(); err != nil {
			halt("Render", err.Error())
		}
		if now := time.Now(); !now.Before(nextHeartbeat) {
			println("zero-kb02 24: heartbeat, scrolling")
			nextHeartbeat = now.Add(2 * time.Second)
		}
		time.Sleep(40 * time.Millisecond)
	}
}

func halt(stage, detail string) {
	for {
		println("zero-kb02 24: FAILED at", stage, ":", detail)
		time.Sleep(2 * time.Second)
	}
}

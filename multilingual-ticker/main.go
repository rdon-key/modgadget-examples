//go:build tinygo

package main

import (
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/efont16"
	board "github.com/rdon-key/modgadget/device/cardputeradv"
)

const message = "<style=english>English </style>" +
	"<style=japanese>日本語 </style>" +
	"<style=chinese>中文 </style>" +
	"<style=korean><b>한국어</b></style>"

func main() {
	time.Sleep(3 * time.Second)
	panel, err := board.ConfigureDisplay()
	if err != nil {
		panic(err)
	}
	font := efont16.Font
	gadget := modgadget.New(panel, modgadget.WithStyles(modgadget.StyleSet{
		Default: modgadget.Style{Font: font, Foreground: modgadget.ColorWhite, Background: modgadget.ColorBlack},
		Entries: []modgadget.StyleEntry{
			{Name: "english", Style: modgadget.Style{Font: font, Foreground: modgadget.RGB565(255, 64, 64), Background: modgadget.ColorBlack}},
			{Name: "japanese", Style: modgadget.Style{Font: font, Foreground: modgadget.RGB565(255, 208, 64), Background: modgadget.ColorBlack}},
			{Name: "chinese", Style: modgadget.Style{Font: font, Foreground: modgadget.RGB565(96, 224, 128), Background: modgadget.ColorBlack}},
			{Name: "korean", Style: modgadget.Style{Font: font, Foreground: modgadget.RGB565(96, 224, 128), Background: modgadget.ColorBlack}},
		},
	}))
	if err := gadget.Clear(); err != nil {
		panic(err)
	}
	view := gadget.Viewport(modgadget.Bounds(0, 0, board.DisplayWidth, 16))
	if err := view.SetText(message); err != nil {
		panic(err)
	}
	view.SetHorizontalScroll(
		modgadget.ScrollSpeed(24),
		modgadget.ScrollGap(32),
		modgadget.ScrollLoop(),
		modgadget.ScrollFromRight(),
	)

	for {
		gadget.Update(time.Now())
		if err := gadget.Render(); err != nil {
			panic(err)
		}
		time.Sleep(16 * time.Millisecond)
	}
}

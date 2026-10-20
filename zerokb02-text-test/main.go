//go:build tinygo

package main

import (
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/spleen8x16"
	board "github.com/rdon-key/modgadget/device/zerokb02"
)

// Small-font text-rendering diagnostic for Zero-KB02.
// Separates the ModGadget layout/raster pipeline from large Efont16 startup.
func main() {
	time.Sleep(2 * time.Second)

	panel, err := board.ConfigureDisplay()
	if err != nil {
		for {
			println("zero-kb02 text test: ConfigureDisplay:", err.Error())
			time.Sleep(2 * time.Second)
		}
	}

	font := spleen8x16.Font
	gadget := modgadget.New(panel, modgadget.WithStyles(modgadget.StyleSet{
		Default: modgadget.Style{
			Font:       font,
			Foreground: modgadget.ColorWhite,
			Background: modgadget.ColorBlack,
		},
	}))

	stage := "OLED configured"
	if err := gadget.Clear(); err != nil {
		stage = "clear error: " + err.Error()
	} else {
		view := gadget.Viewport(modgadget.Bounds(12, 20, 104, 24))
		if err := view.SetText("HELLO"); err != nil {
			stage = "SetText error: " + err.Error()
		} else if err := gadget.Render(); err != nil {
			stage = "Render error: " + err.Error()
		} else {
			stage = "HELLO rendered"
		}
	}

	for {
		println("zero-kb02 text test:", stage)
		time.Sleep(2 * time.Second)
	}
}

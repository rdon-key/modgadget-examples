//go:build tinygo

package main

import (
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/efont16"
	board "github.com/rdon-key/modgadget/device/cardputeradv"
)

const message = "English 日本語 中文 한국어"

func main() {
	time.Sleep(3 * time.Second)
	panel, err := board.ConfigureDisplay()
	if err != nil {
		panic(err)
	}
	gadget := modgadget.New(panel, modgadget.WithStyles(modgadget.StyleSet{
		Default: modgadget.Style{Font: efont16.Font, Foreground: modgadget.ColorWhite, Background: modgadget.ColorBlack},
	}))
	if err := gadget.Clear(); err != nil {
		panic(err)
	}
	view := gadget.Viewport(modgadget.Bounds(0, 0, board.DisplayWidth, 16))
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

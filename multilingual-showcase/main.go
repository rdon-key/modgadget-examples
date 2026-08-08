//go:build tinygo

package main

import (
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/efont24"
	board "github.com/rdon-key/modgadget/device/cardputeradv"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	panel, err := board.ConfigureDisplay()
	must(err)

	white := modgadget.Style{
		Font:       efont24.Font,
		Foreground: modgadget.ColorWhite,
		Background: modgadget.ColorBlack,
	}

	cyan := modgadget.Style{
		Font:       efont24.Font,
		Foreground: modgadget.RGB565(80, 200, 255),
		Background: modgadget.ColorBlack,
	}

	gadget := modgadget.New(
		panel,
		modgadget.WithStyles(modgadget.StyleSet{
			Default: white,
			Entries: []modgadget.StyleEntry{
				{Name: "white", Style: white},
				{Name: "cyan", Style: cyan},
			},
		}),
	)

	must(gadget.Clear())

	width, height := panel.Size()

	view := gadget.Viewport(
		modgadget.Bounds(0, (height-24)/2, width, 24),
	)

	view.SetHorizontalScroll(
		modgadget.ScrollSpeed(50),
		modgadget.ScrollFromRight(),
	)

	must(view.SetText(
		"<style=white>Hello, English</style> " +
			"<style=cyan>こんにちは、日本</style> " +
			"<style=white>你好，中文</style> " +
			"<style=cyan>你好，臺灣</style> " +
			"<style=white>안녕하세요, 한국어</style> " +
			"<style=cyan>Olá, Português</style> " +
			"<style=white>Pozdrav, Hrvatski</style> " +
			"<style=cyan>Привет, Русский</style> " +
			"<style=white>Γεια σου, Ελληνικά</style>",
	))

	for {
		gadget.Update(time.Now())
		must(gadget.Render())
		time.Sleep(16 * time.Millisecond)
	}
}

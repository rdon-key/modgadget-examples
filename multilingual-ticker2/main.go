//go:build tinygo

package main

import (
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/efont16"
	"github.com/rdon-key/modgadget-fonts/efont24"
	board "github.com/rdon-key/modgadget/device/cardputeradv"
)

const (
	styleEnglish  = "english"
	styleJapanese = "japanese"
	styleChinese  = "chinese"
	styleKorean   = "korean"

	timerY      = 111
	timerHeight = 16
	maximumTime = 99*60 + 59
)

type tickerConfig struct {
	text          string
	x, y          int16
	width, height int16
	speed         float64
	gap           int16
}

var tickers = [...]tickerConfig{
	{text: "<style=english>English</style>", x: 0, y: 0, width: 240, height: 24, speed: 20, gap: 40},
	{text: "<style=japanese>日本語</style>", x: 0, y: 27, width: 240, height: 24, speed: 24, gap: 48},
	{text: "<style=chinese>中文</style>", x: 0, y: 54, width: 240, height: 24, speed: 60, gap: 56},
	{text: "<style=korean>한국어</style>", x: 0, y: 81, width: 240, height: 24, speed: 50, gap: 64},
}

func formatElapsed(buffer *[5]byte, totalSeconds uint64) string {
	if totalSeconds > maximumTime {
		totalSeconds = maximumTime
	}
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	buffer[0] = byte('0' + minutes/10)
	buffer[1] = byte('0' + minutes%10)
	buffer[2] = ':'
	buffer[3] = byte('0' + seconds/10)
	buffer[4] = byte('0' + seconds%10)
	return string(buffer[:])
}

func main() {
	time.Sleep(3 * time.Second)
	panel, err := board.ConfigureDisplay()
	if err != nil {
		panic(err)
	}

	font24 := efont24.Font
	styles := modgadget.StyleSet{
		Default: modgadget.Style{Font: efont16.Font, Foreground: modgadget.ColorWhite, Background: modgadget.ColorBlack},
		Entries: []modgadget.StyleEntry{
			{Name: styleEnglish, Style: modgadget.Style{Font: font24, Foreground: modgadget.RGB565(255, 64, 64), Background: modgadget.ColorBlack}},
			{Name: styleJapanese, Style: modgadget.Style{Font: font24, Foreground: modgadget.RGB565(255, 208, 64), Background: modgadget.ColorBlack}},
			{Name: styleChinese, Style: modgadget.Style{Font: font24, Foreground: modgadget.RGB565(96, 224, 128), Background: modgadget.ColorBlack}},
			{Name: styleKorean, Style: modgadget.Style{Font: font24, Foreground: modgadget.RGB565(64, 192, 255), Background: modgadget.ColorBlack}},
		},
	}
	gadget := modgadget.New(panel, modgadget.WithStyles(styles))
	if err := gadget.Clear(); err != nil {
		panic(err)
	}

	for _, config := range tickers {
		view := gadget.Viewport(modgadget.Bounds(config.x, config.y, config.width, config.height))
		if err := view.SetText(config.text); err != nil {
			panic(err)
		}
		view.SetHorizontalScroll(
			modgadget.ScrollSpeed(config.speed),
			modgadget.ScrollGap(config.gap),
			modgadget.ScrollLoop(),
			modgadget.ScrollFromRight(),
		)
	}

	timerMeasurement, err := modgadget.MeasureText("00:00", styles)
	if err != nil {
		panic(err)
	}
	timerX := board.DisplayWidth - timerMeasurement.Width
	timer := gadget.Viewport(modgadget.Bounds(timerX, timerY, timerMeasurement.Width, timerHeight))

	start := time.Now()
	lastSecond := uint64(maximumTime + 1)
	var timerBuffer [5]byte
	for {
		now := time.Now()
		elapsed := uint64(now.Sub(start) / time.Second)
		if elapsed > maximumTime {
			elapsed = maximumTime
		}
		if elapsed != lastSecond {
			if err := timer.SetText(formatElapsed(&timerBuffer, elapsed)); err != nil {
				panic(err)
			}
			lastSecond = elapsed
		}

		gadget.Update(now)
		if err := gadget.Render(); err != nil {
			panic(err)
		}
		time.Sleep(16 * time.Millisecond)
	}
}

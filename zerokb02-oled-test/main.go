//go:build tinygo

package main

import (
	"time"

	board "github.com/rdon-key/modgadget/device/zerokb02"
)

// This test deliberately imports no font packages: it tests OLED setup,
// I2C transfers and RGB565-to-monochrome conversion independently of MGZ.
func main() {
	time.Sleep(2 * time.Second)
	println("zero-kb02: starting OLED diagnostics")

	display, err := board.ConfigureDisplay()
	if err != nil {
		println("zero-kb02: ConfigureDisplay failed:", err.Error())
		for {
			time.Sleep(time.Second)
		}
	}
	println("zero-kb02: OLED configured")

	var row [128 * 2]byte
	white := true
	for {
		var value byte
		if white {
			value = 0xff
		}
		for i := range row {
			row[i] = value
		}

		if err := display.BeginRect(0, 0, board.DisplayWidth, board.DisplayHeight); err != nil {
			println("zero-kb02: BeginRect failed:", err.Error())
			break
		}
		writeOK := true
		for y := int16(0); y < board.DisplayHeight; y++ {
			if err := display.WritePixels(row[:]); err != nil {
				println("zero-kb02: WritePixels failed:", err.Error())
				writeOK = false
				break
			}
		}
		if !writeOK {
			break
		}
		if err := display.EndRect(); err != nil {
			println("zero-kb02: EndRect failed:", err.Error())
			break
		}
		if white {
			println("zero-kb02: white frame sent")
		} else {
			println("zero-kb02: black frame sent")
		}
		white = !white
		time.Sleep(2 * time.Second)
	}
	for {
		time.Sleep(time.Second)
	}
}

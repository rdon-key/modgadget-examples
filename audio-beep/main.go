//go:build tinygo

package main

import (
	"runtime"
	"time"

	audio "github.com/rdon-key/modgadget/audio/cardputeradv"
	board "github.com/rdon-key/modgadget/device/cardputeradv"
)

var sounds = [...]audio.Pattern{
	audio.PatternStartup,
	audio.PatternClick,
	audio.PatternCorrect,
}

func play(player *audio.Player, pattern audio.Pattern) {
	if err := player.PlayPattern(pattern); err != nil {
		panic(err)
	}
	for player.Busy() {
		if err := player.Update(); err != nil {
			_ = player.Stop()
			panic(err)
		}
		runtime.Gosched()
	}
}

func main() {
	player, err := board.ConfigureAudio()
	if err != nil {
		panic(err)
	}
	for _, pattern := range sounds {
		play(player, pattern)
		time.Sleep(300 * time.Millisecond)
	}
	for {
		time.Sleep(time.Second)
	}
}

# audio-beep

This is the standalone example of ModGadget's public Cardputer ADV audio API.
It configures the device with `device/cardputeradv.ConfigureAudio`, then plays
the built-in `PatternStartup`, `PatternClick`, and `PatternCorrect` sounds once.
It calls `Player.Update` while each sound is busy and remains silent after the
sequence completes.

- Hardware: M5Stack Cardputer ADV
- Main APIs: `device/cardputeradv.ConfigureAudio`; `audio/cardputeradv.PatternStartup`, `PatternClick`, `PatternCorrect`, and `Player.PlayPattern`, `Update`, `Busy`, `Stop`
- Font: none
- Expected sound: Startup, Click, and Correct once each, separated by 300 ms
- Controls: none; the example does not use the display or keyboard

```sh
tinygo build -target=m5stamp-s3a ./audio-beep
tinygo flash -target=m5stamp-s3a -port=<PORT> ./audio-beep
```

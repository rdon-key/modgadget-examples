# audio-beep

This minimal Cardputer ADV audio example plays three built-in sound patterns
once. It calls `Update` on the cooperative player while waiting for each sound
to finish, then remains silent after the sequence completes.

- Hardware: M5Stack Cardputer ADV
- Public API: `device/cardputeradv.ConfigureAudio`, `audio/cardputeradv.Player.PlayPattern`, `Update`, `Busy`, `Stop`
- Font: none
- Expected sound: Startup, Click, and Correct once each, separated by 300 ms
- Controls: none; the example does not use the display or keyboard

```sh
tinygo build -target=m5stamp-s3a ./audio-beep
tinygo flash -target=m5stamp-s3a -port=<PORT> ./audio-beep
```

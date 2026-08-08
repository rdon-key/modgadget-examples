# keyboard-typing

Small Cardputer ADV keyboard example using `WithKeyboard`, `OnKey`, `KeyDown`,
printable `KeyEvent.Rune`, and Backspace to update a typed-text `Viewport`.
The audio controller is registered so Fn volume and mute shortcuts are consumed
by ModGadget rather than becoming typed input.
Typed text is escaped before it is passed to markup, so `<` is displayed
literally and does not start a markup tag.

It uses ModGadget and the Efont 16 package from `modgadget-fonts` as normal Go
module dependencies.

- Hardware: M5Stack Cardputer ADV
- Public API: `WithKeyboard`, `WithVolumeController`, `OnKey`, `KeyEvent`, `Viewport`, `SetText`, `Update`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Expected display: a title, the typed text, and Backspace help
- Controls: type with the character keys; Backspace removes one character

Run the host-side state tests with `go test ./keyboard-typing` from the module
root. Build the hardware program with TinyGo:

```sh
tinygo build -target=m5stamp-s3a ./keyboard-typing
tinygo flash -target=m5stamp-s3a -port=<PORT> ./keyboard-typing
```

Flashing is a separate operation and is not part of the build.

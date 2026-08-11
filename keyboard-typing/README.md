# keyboard-typing

Small Cardputer ADV keyboard example that adds printable characters on key-down
events and uses Backspace to remove one character from a typed-text `Viewport`.
Cardputer ADV Fn volume and mute shortcuts adjust audio volume and are not inserted
as text.
Typed text is escaped before it is passed to markup, so `<` is displayed
literally and does not start a markup tag.

It uses ModGadget and the Efont 16 package from `modgadget-fonts` as normal Go
module dependencies.

- Hardware: M5Stack Cardputer ADV
- Main APIs: `device/cardputeradv.ConfigureKeyboard`, `ConfigureAudio`; `WithKeyboard`, `WithVolumeController`, `OnKey`, `KeyEvent`, `KeyDown`, `KeyBackspace`, `Viewport`, `SetText`, `Update`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Expected display: a title, the typed text, and Backspace help
- Controls: type with the character keys; Backspace removes one character; Cardputer ADV Fn volume and mute shortcuts control audio volume

Run the host-side state tests with `go test ./keyboard-typing` from the module
root. Build the hardware program with TinyGo:

```sh
tinygo build -target=m5stamp-s3a ./keyboard-typing
tinygo flash -target=m5stamp-s3a -port=<PORT> ./keyboard-typing
```

Flashing is a separate operation and is not part of the build.

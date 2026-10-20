# Zero-KB02: Traditional Chinese with a compressed font

Display **恭喜進軍台灣！** ("Congratulations on expanding to Taiwan!")
on one line in a 16-dot bitmap font on the TinyGo Keeb Tour Zero-KB02's 128x64 monochrome OLED.

This example consumes ModGadget and its ready-to-use compressed
`efont16.Font` from [modgadget-fonts](https://github.com/rdon-key/modgadget-fonts).
It does not generate or copy fonts into ModGadget itself. The OLED adapter is
provided by `github.com/rdon-key/modgadget/device/zerokb02`.

## Hardware

- Zero-KB02 based on the Waveshare RP2040-Zero (2 MB flash)
- SSD1306 OLED: 128x64, I2C0, GPIO12 SDA, GPIO13 SCL, address 0x3c
- Physical OLED orientation: 180 degrees

If the flash operation succeeds but the screen remains blank, first run
[`zerokb02-oled-test`](../zerokb02-oled-test/) to separate SSD1306/I2C
problems from font/text rendering. Check the greeting glyphs with
`go test ./zerokb02-traditional -run TestGreetingGlyphCoverage -v`.

### Serial diagnostics

The OLED white/black diagnostics and the host glyph coverage test have
both been confirmed working, but the full greeting initially remained blank.
The greeting example now emits progress information on USB CDC serial.

```sh
tinygo flash -target=waveshare-rp2040-zero ./zerokb02-traditional
tinygo monitor -target=waveshare-rp2040-zero
```

Look for `greeting demo starting`, `OLED configured`, the seven
`lookup glyph` messages, `all glyphs decoded`, `OLED cleared`,
`building text layout`, `rendering text`, and `render complete`.
The last printed stage identifies where execution stops. If `lookup glyph`
fails on the device but passes in the host test, investigate the MGZ
DEFLATE decoding path on TinyGo. If no startup text appears even after
resetting with the monitor connected, initialization of the imported font
package may have panicked before `main` began.

The example currently demonstrates **display only**. Keyboard, joystick,
rotary encoder, LEDs and USB HID are not initialized.

## Build and flash

From the `modgadget-examples` repository root:

```sh
go mod download

# Until ModGadget PR #13 is merged, use the commit that adds Zero-KB02:
go get github.com/rdon-key/modgadget@d0b4f8a90fb4edbcf69116e405498f241315355b

tinygo build -target=waveshare-rp2040-zero -size=short -o zerokb02.uf2 ./zerokb02-traditional
tinygo flash -target=waveshare-rp2040-zero ./zerokb02-traditional
```

The `go get` line temporarily updates the *local copy* of `go.mod` and
`go.sum` to a ModGadget revision that includes `device/zerokb02`.
It is needed while [ModGadget PR #13](https://github.com/rdon-key/modgadget/pull/13)
is unmerged; when the driver becomes available in the normal ModGadget
dependency, use that version instead.

The existing `modgadget-fonts` dependency already provides `efont16`.
No extra `go get` of the font package is necessary.

**Flash budget:** the compressed full Efont16 asset is 876,150 bytes,
639,512 bytes smaller than Efont24 (1,515,662 bytes). A previous
Efont24 build exceeded the 2 MB RP2040-Zero flash by 567,252 bytes.
A user confirmed a flashable 16-dot build at 970,660 bytes of flash,
but the OLED remained blank; display diagnostics are pending.
If it is still too large, generate a subset MGZ using
[modgadget-font-assets](https://github.com/rdon-key/modgadget-font-assets)
and distribute it as a font package.

The 16-dot firmware builds and flashes on Zero-KB02; successful OLED output
has **not** yet been verified.

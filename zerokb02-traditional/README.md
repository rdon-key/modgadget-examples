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
The 16-dot build is expected to fit, but verify the actual link result.
If it is still too large, generate a subset MGZ using
[modgadget-font-assets](https://github.com/rdon-key/modgadget-font-assets)
and distribute it as a font package.

This example has not yet been cross-built or verified on a physical Zero-KB02.

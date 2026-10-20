# Zero-KB02: Traditional Chinese with a compressed font

Display **恭喜進軍台灣！** ("Congratulations on expanding to Taiwan!")
on one line in a 16-dot bitmap font on the TinyGo Keeb Tour Zero-KB02's 128x64 monochrome OLED.

This example consumes ModGadget and its ready-to-use compressed
`efont16zerokb02.Font` from [modgadget-fonts](https://github.com/rdon-key/modgadget-fonts/pull/1).
This is a 17,001-byte compressed subset of Efont16 with 448 glyphs,
including all seven characters in the greeting.
It does not generate or copy fonts into ModGadget itself. The OLED adapter is
provided by `github.com/rdon-key/modgadget/device/zerokb02`.

## Hardware

- Zero-KB02 based on the Waveshare RP2040-Zero (2 MB flash)
- SSD1306 OLED: 128x64, I2C0, GPIO12 SDA, GPIO13 SCL, address 0x3c
- Physical OLED orientation: 180 degrees

If the flash operation succeeds but the screen remains blank, first run
[`zerokb02-oled-test`](../zerokb02-oled-test/) to separate SSD1306/I2C
problems from font/text rendering. If the white/black OLED test passes, run
[`zerokb02-text-test`](../zerokb02-text-test/) to render `HELLO` with a
small 8x16 MGZ font through the same ModGadget viewport API. Check the greeting
glyphs with `go test ./zerokb02-traditional -run TestGreetingGlyphCoverage -v`.

### RP2040 out-of-memory diagnosis

The earlier 17KB font successfully opened, but the RP2040 logged
`fatal error: out of memory` while looking up the sixth glyph (`灣`).
The seven greeting glyphs occupy seven different MGZ DEFLATE blocks.
`mgz.Font` previously created a new `flate.NewReader` for each block,
which can consume substantially more heap than the compressed asset size.

[ModGadget PR #14](https://github.com/rdon-key/modgadget/pull/14)
reuses a single DEFLATE decoder with `flate.Resetter.Reset` across cache
misses. The revised ModGadget revision above includes **both** the OLED
driver (PR #13) and this tentative RAM improvement (PR #14).

The reader-reuse fix is not yet verified on hardware. A TinyGo GC bug
has **not** been established; the current code's allocation pattern
itself can exhaust a constrained MCU heap.

### Startup and OLED diagnostics

The full 16-dot font and the first 17 KB subset both built and flashed
without errors, yet neither produced any serial heartbeat. In contrast,
the OLED black/white diagnostic and the smaller Spleen font's `HELLO`
text rendering both worked. The original Efont16 package and the first
subset package both opened the font **before** `main()`, via a global
`MustOpenMGZ` initializer.

The subset package now provides `Open()`, allowing the application to
enter `main()`, initialize the OLED, and only then parse the MGZ file.

On startup, before attempting to open the font, the greeting example:
1. Draws a **solid white 16x8 marker at the top-left** of the OLED.
2. Prints a countdown to opening the font for 8 seconds.
3. Calls `efont16zerokb02.Open()` and attempts to render the greeting.

**If the marker never appears**, the firmware may fail before or
during `ConfigureDisplay()`. **If the marker appears but no greeting
follows**, the font-opening and later rendering stages are suspects.
If the greeting succeeds, the marker is removed by `gadget.Clear()`.
Connect the serial monitor immediately after flashing to see the
countdown, then check for the repeated final status message.

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
After rendering, the greeting demo repeats a `heartbeat, render complete`
message every 2 seconds. All handled failure paths repeat a `FAILED at ...`
message as well, even when the monitor connects late. If `lookup glyph`
fails on the device but passes in the host test, investigate the MGZ
DEFLATE decoding path on TinyGo. If **no messages appear at all** with the
repeating logs, the firmware may have stopped before `main` (for example,
during initialization of the imported full Efont16 package).

The example currently demonstrates **display only**. Keyboard, joystick,
rotary encoder, LEDs and USB HID are not initialized.

## Build and flash

From the `modgadget-examples` repository root:

```sh
go mod download

# Until ModGadget PRs #13 and #14 merge, use the OLED driver + MGZ inflater fix:
go get github.com/rdon-key/modgadget@680edd09ff0dce5daedd7df0a03f0697dba29f71

# Also use the standalone MGZ subset in modgadget-fonts PR #1:
go get github.com/rdon-key/modgadget-fonts@885d3a279f8b8b2194097e3f6e907068481f063f

tinygo build -target=waveshare-rp2040-zero -size=short -o zerokb02.uf2 ./zerokb02-traditional
tinygo flash -target=waveshare-rp2040-zero ./zerokb02-traditional
```

The `go get` line temporarily updates the *local copy* of `go.mod` and
`go.sum` to a ModGadget revision that includes `device/zerokb02`.
It is needed while [ModGadget PR #13](https://github.com/rdon-key/modgadget/pull/13)
is unmerged; when the driver becomes available in the normal ModGadget
dependency, use that version instead.

The repository already depends on `modgadget-fonts`, but the original
revision does not contain `efont16zerokb02`. The second `go get` above
upgrades that dependency to a commit containing the new compressed subset.
Both commands temporarily update the local `go.mod` and `go.sum`.

**Flash budget and startup:** the full Efont24 MGZ is 1,515,662 bytes,
full Efont16 MGZ is 876,150 bytes, and this compressed Efont16 subset
is **17,001 bytes**. The full Efont16 firmware fitted in flash (970,660
bytes) but produced no serial heartbeat and left the OLED blank.
Both the font-free OLED test and small-font ModGadget `HELLO` test worked.
The subset greatly reduces MGZ metadata validation at the time of an
explicit `Open()` call and leaves ample flash room. The earlier subset
with global initialization also failed to start; the deferred-open variant
still requires real-device testing.

The full Efont16 firmware built and flashed but the OLED was blank.
The first 17KB subset flashed but produced no serial output. The new
explicit-open variant has not yet been tested on Zero-KB02.

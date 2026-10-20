# Zero-KB02: Traditional Chinese with full Efont16

Display **恭喜進軍台灣！** on the TinyGo Keeb Tour Zero-KB02's
128x64 SSD1306 OLED, using the **complete Efont Biwidth 16** compressed
bitmap font rather than a demo-specific subset.

The full MGZ file is 876,150 bytes and includes **24,618 glyphs**.
"Complete" means all glyphs in the distributed Efont16 font, not
every Unicode code point. This font already exists in the regular
[`modgadget-fonts/efont16`](https://github.com/rdon-key/modgadget-fonts/tree/main/efont16)
package. No font files are copied into ModGadget or this example repo,
and the experimental 17 KB subset PR is no longer required for the example.

The OLED adapter is provided by
[`modgadget/device/zerokb02`](https://github.com/rdon-key/modgadget/pull/13).

## Hardware

- TinyGo Keeb Tour Zero-KB02 based on Waveshare RP2040-Zero
- 128x64 SSD1306 monochrome OLED
- I2C0, GPIO12 SDA, GPIO13 SCL, address `0x3c`
- Display rotation: 180 degrees (to account for physical mounting)

## Build and flash

From the `modgadget-examples` repository root:

```sh
git pull origin feat/zerokb02-traditional-demo

# Use the rebased Zero-KB02 OLED driver from ModGadget PR #13.
# This commit also includes the already-merged MGZ decoder fix (PR #14).
go get github.com/rdon-key/modgadget@6130268fd035a5ab659c6792aaf8cd050cf07a1a

# The full Efont16 font is already in the existing modgadget-fonts dependency.
go test ./zerokb02-traditional -run TestGreetingGlyphCoverage -v

# Show flash/RAM sizes and attach USB CDC monitor immediately.
tinygo flash -target=waveshare-rp2040-zero -size=short -monitor ./zerokb02-traditional
```

**Flash capacity:** The full Efont16 MGZ embeds 876,150 bytes, compared
with 17,001 bytes for the prior diagnostic subset. An earlier full Efont16
build occupied 970,660 bytes of flash; the exact size of this revised
firmware still needs to be measured. The
`waveshare-rp2040-zero` target's linker reserves a 1020 KiB firmware
region, so flash headroom is limited. Check the new `-size=short`
result instead of assuming the build will always fit. The reported
`ram` value includes static data and BSS; it is **not** peak heap usage.

## OLED and MGZ diagnostics

On startup, the application:

1. Configures the OLED and draws a 16x8 white square in the top-left.
2. Prints a countdown for eight seconds, giving time to attach USB CDC.
3. Checks that all seven characters of the greeting are available from the
   **full** Efont16 font.
4. Clears the OLED and draws the greeting in a centered viewport.
5. Prints `heartbeat, render complete` every two seconds.

The `efont16.Font` value is opened during Go package initialization,
**before `main()` runs**. This is different from the previous diagnostic
subset's explicit `Open()` call. If the application emits no logs or
square, investigate package startup, static initialization, and memory.

If display hardware behavior is uncertain, use
[`zerokb02-oled-test`](../zerokb02-oled-test/) to toggle all-white and
all-black frames. Use
[`zerokb02-text-test`](../zerokb02-text-test/) to render `HELLO` with
a small Spleen MGZ font through the same ModGadget text API.

## MGZ decompression and RAM

The first full-font firmware was black and silent. During subsequent
diagnostics with the small Efont16 subset, the device revealed
`fatal error: out of memory` on the sixth different compressed glyph
block. [ModGadget PR #14](https://github.com/rdon-key/modgadget/pull/14)
fixed this by reusing a single `compress/flate` decoder across
consecutive block cache misses. The 17 KB subset then rendered the
greeting successfully on hardware, and PR #13 was rebased and retested
successfully on the same board.

The **full-font build using the decoder reuse fix has not yet been
verified on real hardware**. Until tested, do not assume its 24,618-glyph
index or startup initialization behaves identically to the subset.

This example demonstrates **display output only**. Keyboard, joystick,
rotary encoder, LEDs and USB HID are not initialized.

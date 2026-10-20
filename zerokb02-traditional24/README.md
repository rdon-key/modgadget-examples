# Zero-KB02: full Efont24 horizontal scroll

Display the complete Traditional Chinese greeting **恭喜進軍台灣！**
using the **full Efont Biwidth 24** compressed MGZ font on the 128×64
SSD1306 OLED. The seven characters are rendered on **one line** and
scroll continuously from right to left; there is no line wrapping.

The existing `github.com/rdon-key/modgadget-fonts/efont24` package
provides the complete font (MGZ size: **1,515,662 bytes**) and is
embedded in firmware flash, not copied into RAM as a whole. The
embedded font's index is parsed during package initialization. MGZ
blocks are decompressed on demand using the reusable DEFLATE reader
from [ModGadget PR #14](https://github.com/rdon-key/modgadget/pull/14).

## Why a custom target?

The hardware is a **Waveshare RP2040-Zero with 2 MiB of onboard flash**.
However, TinyGo's built-in `waveshare-rp2040-zero` target sets
`--defsym=__flash_size=1020K`, limiting linked firmware to
1,044,480 bytes. The full Efont24 MGZ is **already larger than that
limit on its own**. The earlier failure to build Efont24 therefore
does **not** show that the board's 2 MiB flash is inadequate.

[`../targets/zerokb02-2mb.json`](../targets/zerokb02-2mb.json)
inherits TinyGo's `rp2040` platform and duplicates the board-specific
settings from `waveshare-rp2040-zero`, with only the firmware flash
limit changed to **2048K**. It retains the Waveshare build tag, USB
serial VID/PID, BOOT2 image and stack size.

**Safety:** Use this custom target only on a board confirmed to have
**2 MiB of physical flash**. The linker setting does not increase
physical flash. Flashing a different 1 MiB board with the larger
firmware can fail or corrupt other data. Even on the correct board,
flashing may overwrite any user data previously stored in the upper
flash region. The built-in 1020K limit may reflect conservative
partitioning or other target choices, so use this override deliberately.

## Build, inspect size and flash

Run from the `modgadget-examples` repository root:

```sh
git pull origin feat/zerokb02-traditional-demo

# PR #13 contains the Zero-KB02 display adapter and the merged
# decoder-reuse fix (PR #14). Keep PR #13 unmerged for this test.
go get github.com/rdon-key/modgadget@6130268fd035a5ab659c6792aaf8cd050cf07a1a

go test -count=1 ./zerokb02-traditional24 -v

# Compile first so you can inspect the resulting firmware size.
tinygo build -target=./targets/zerokb02-2mb.json \
  -size=short -o zerokb02-24.uf2 ./zerokb02-traditional24

# Only flash if the build succeeds and the physical flash size is 2 MiB.
tinygo flash -target=./targets/zerokb02-2mb.json \
  -size=short -monitor ./zerokb02-traditional24
```

The `flash` size must be **less than 2,097,152 bytes**. `ram` as
reported by `-size=short` is static data + BSS, not peak runtime
heap usage. This program has not yet been built or flashed on the
hardware with the new target; record both the build size and the
USB serial log when testing it.

### Expected behavior

- The full Efont24 font opens successfully at firmware startup.
- USB prints `zero-kb02 24: starting full Efont24 horizontal scroll`
  followed by glyph lookups, `all seven glyphs decoded` and
  `starting horizontal marquee`.
- Text scrolls from right to left at **24 pixels/second**.
- USB prints `heartbeat, scrolling` every two seconds.
- If an operation returns an error, a repeated `FAILED at ...`
  message identifies the stage.

If the USB drive `RPI-RP2` cannot be found, that is a **USB/bootloader
enumeration** problem, not a linker flash-overflow diagnosis. A prior
instance was resolved by restarting the Windows PC.

## Other available diagnostics

The working 16-dot full-font example is intentionally retained as
[`zerokb02-traditional`](../zerokb02-traditional/). If this experiment
fails, flash it again using the normal `waveshare-rp2040-zero`
target and its documented command.

For lower-level checks, use
[`zerokb02-oled-test`](../zerokb02-oled-test/)
(white/black panel) or
[`zerokb02-text-test`](../zerokb02-text-test/)
(small-font `HELLO` rendering).

This demo covers the OLED only; keyboard, joystick, encoder, LEDs
and USB HID are not configured.

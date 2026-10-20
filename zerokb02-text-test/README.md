# Zero-KB02 text rendering diagnostic (small MGZ font)

Shows **HELLO** on the Zero-KB02 monochrome OLED with the standalone
Spleen 8x16 font package (only about 20 KB compressed).

This uses the *same* ModGadget `Viewport`, `SetText` and `Render`
pipeline as `zerokb02-traditional`, but does not import the large
Efont16 font. Use it if `zerokb02-oled-test` alternates white and black
successfully while the Traditional Chinese greeting remains blank.

```sh
tinygo flash -target=waveshare-rp2040-zero ./zerokb02-text-test
tinygo monitor -target=waveshare-rp2040-zero
```

The monitor prints a status **every two seconds**, so it is safe to
connect after flashing. A successful test shows `HELLO` on screen and
prints `HELLO rendered` repeatedly. Errors are also printed repeatedly.

- `HELLO` visible: the OLED, ModGadget viewport pipeline and small
  MGZ decoding work; investigate Efont16 initialization or its much
  larger compressed-font decoding path.
- Still blank, `HELLO rendered`: investigate viewport/raster output on
  the monochrome display.
- No serial output at all: program may not have entered `main`; compare
  with the known working `zerokb02-oled-test`.

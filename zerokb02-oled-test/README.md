# Zero-KB02 OLED diagnostics (no fonts)

This minimal hardware test distinguishes an OLED / I2C issue from font /
ModGadget text rendering.

It uses only `device/zerokb02.ConfigureDisplay` and the public ModGadget
display rectangle streaming API; it does **not** load Efont16 or MGZ.

```sh
# In modgadget-examples, first make sure you use a ModGadget revision
# containing device/zerokb02; see the zerokb02-traditional example README.
tinygo flash -target=waveshare-rp2040-zero ./zerokb02-oled-test
tinygo monitor -target=waveshare-rp2040-zero
```

The screen should alternate **fully white** and **fully black** every two
seconds. The serial monitor prints `OLED configured`, `white frame sent`,
or an error (if an I2C call fails). Reflash while the monitor is connected if
you missed startup output.

How to interpret the results:

- White / black alternating: I2C and the basic SSD1306 backend work; test
  font glyph coverage and text drawing next.
- `ConfigureDisplay failed`: the OLED initialization or I2C bus failed.
- `EndRect failed`: the display data write failed.
- `white frame sent` repeatedly but panel stays dark: display might not
  accept our initialization/addressing sequence, or hardware requires separate
  checks against the known working workshop `07_oled` example.

For the hardware wiring and known-good example see
[the TinyGo Keeb workshop](https://github.com/tinygo-keeb/workshop).

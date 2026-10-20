# ModGadget Examples

This repository contains standalone TinyGo examples using [ModGadget](https://github.com/rdon-key/modgadget). Most examples target the M5Stack Cardputer ADV; the `zerokb02-traditional` example targets the TinyGo Keeb Tour Zero-KB02.

ModGadget and [ModGadget Fonts](https://github.com/rdon-key/modgadget-fonts) are obtained as ordinary Go module dependencies. A local clone of the ModGadget repository is not required.

Each example can be built independently and demonstrates a specific part of the ModGadget application model, ranging from static multilingual text to keyboard-driven applications.

## Examples

| Example | Description |
| --- | --- |
| [`multilingual-text`](multilingual-text/) | Static English, Japanese, Chinese, and Korean text |
| [`multilingual-ticker`](multilingual-ticker/) | Looping horizontal scroll with four languages |
| [`multilingual-ticker2`](multilingual-ticker2/) | Independent 24-dot multilingual tickers with different positions and speeds, plus an elapsed-time display |
| [`multilingual-document`](multilingual-document/) | Application-level vertical document scrolling with a four-language ModGadget introduction |
| [`multilingual-showcase`](multilingual-showcase/) | One-shot 24-dot ticker with nine multilingual greetings, alternating white and cyan |
| [`audio-beep`](audio-beep/) | Standalone example of the public Cardputer ADV audio API, playing three built-in patterns |
| [`keyboard-typing`](keyboard-typing/) | Keyboard-driven text entry with Backspace and Cardputer ADV volume shortcuts |
| [`mgz-text`](mgz-text/) | Static multilingual text using the compressed 24-dot MGZ font package |
| [`zerokb02-traditional`](zerokb02-traditional/) | 24-dot compressed Traditional Chinese text on the Zero-KB02 OLED |

## Complete application

For a complete application built with the same public ModGadget APIs, see:

- [Rdon Type 100](https://github.com/rdon-key/rdon-type100)

Rdon Type 100 combines multilingual fonts, text layout, Viewports, scrolling, keyboard input, audio, and application state into a complete Cardputer ADV typing game.

A prebuilt firmware image is also available from the [Rdon Type 100 Releases](https://github.com/rdon-key/rdon-type100/releases/latest) page.

## Clone

Clone this repository:

```sh
git clone https://github.com/rdon-key/modgadget-examples.git
cd modgadget-examples
```

Download the Go module dependencies:

```sh
go mod download
```

ModGadget and ModGadget Fonts are downloaded automatically through the Go module system.

## TinyGo setup

The Cardputer ADV examples cannot currently be built with a released version of TinyGo. They require TinyGo `0.42.0-dev` built from the TinyGo `dev` branch, which includes support for the M5Stamp-S3A.

The multilingual examples and `keyboard-typing`, which embed multilingual MGZ
font data, also require the ESP32-S3 large-flash fix provided by
[PR #5567](https://github.com/tinygo-org/tinygo/pull/5567). The font-free
`audio-beep` example does not require this fix.

Without this fix, the ESP32-S3 ROM bootloader may reject a large firmware image before TinyGo startup code begins executing.

For complete TinyGo source-build and patch instructions, see the [Rdon Type 100 TinyGo setup](https://github.com/rdon-key/rdon-type100#tinygo-setup).

## Build

Each example can be built independently. The commands below target Cardputer ADV; the Zero-KB02 example uses its own target and dependency instructions in [its README](zerokb02-traditional/README.md).

For example:

```sh
tinygo build -target=m5stamp-s3a ./multilingual-text
```

Other examples can be built by replacing the final directory name:

```sh
tinygo build -target=m5stamp-s3a ./multilingual-ticker
tinygo build -target=m5stamp-s3a ./multilingual-ticker2
tinygo build -target=m5stamp-s3a ./multilingual-document
tinygo build -target=m5stamp-s3a ./multilingual-showcase
tinygo build -target=m5stamp-s3a ./audio-beep
tinygo build -target=m5stamp-s3a ./keyboard-typing
tinygo build -target=m5stamp-s3a ./mgz-text
```

Flashing instructions and example-specific controls are documented in each example's README.

## Test

Logic that can run on the host is tested with the standard Go toolchain:

```sh
go test ./...
```

Hardware-specific `main.go` files use TinyGo build constraints.

Host-side tests do not access embedded hardware. Before flashing an example, also cross-build it with the appropriate TinyGo target. The Zero-KB02 example currently requires the ModGadget Zero-KB02 display branch until [PR #13](https://github.com/rdon-key/modgadget/pull/13) is merged.

## Related projects

| Repository | Purpose |
| --- | --- |
| [modgadget](https://github.com/rdon-key/modgadget) | TinyGo UI/device toolkit and core API-focused or diagnostic examples |
| [modgadget-examples](https://github.com/rdon-key/modgadget-examples) | Standalone practical examples that consume ModGadget as a module dependency |
| [modgadget-fonts](https://github.com/rdon-key/modgadget-fonts) | Importable multilingual MGZ bitmap-font packages |
| [rdon-type100](https://github.com/rdon-key/rdon-type100) | Complete multilingual Cardputer ADV application |
| [modgadget-font-assets](https://github.com/rdon-key/modgadget-font-assets) | Canonical font sources, provenance, generation, and validation |

## License

BSD 3-Clause.

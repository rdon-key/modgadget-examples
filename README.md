# ModGadget Examples

Practical application examples for
[ModGadget](https://github.com/rdon-key/modgadget). These examples combine
multiple library features and complement the small API-focused examples kept
in the ModGadget repository.

The examples use `github.com/rdon-key/modgadget` and
`github.com/rdon-key/modgadget-fonts` as ordinary Go module dependencies. The
module does not require a local checkout of either repository.

## Examples

| example | description | hardware |
| --- | --- | --- |
| [`keyboard-typing`](keyboard-typing/) | Typed-text UI combining direct keyboard events, multiple viewports, markup-safe text, fonts, and ModGadget system audio controls | M5Stack Cardputer ADV |

## Test

Host-side logic and package checks can be run with:

```sh
go test ./...
```

## TinyGo build

The hardware example requires TinyGo and a Cardputer ADV-compatible target.
The current ModGadget setup uses the `m5stamp-s3a` target:

```sh
tinygo build -target=m5stamp-s3a ./keyboard-typing
```

Building does not flash the device. Flashing and serial-port access are
separate operations and are not required to use this repository.

## License

BSD 3-Clause.

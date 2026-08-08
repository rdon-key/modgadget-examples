# multilingual-text

This is the simplest example to try first. It displays English, Japanese,
Chinese, and Korean text in a single static `Viewport`. It does not use
scrolling, keyboard input, or audio.

- Hardware: M5Stack Cardputer ADV
- Public API: `New`, `WithStyles`, `Bounds`, `Viewport`, `SetText`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Expected display: `English 日本語 中文 한국어`
- Controls: none; the text remains on screen after startup

```sh
tinygo build -target=m5stamp-s3a ./multilingual-text
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-text
```

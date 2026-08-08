# multilingual-ticker

This example scrolls `English 日本語 中文 한국어` from right to left as one
looping ticker containing four styles. English is red, Japanese is yellow,
Chinese is green, and Korean is bold green using the public `<b>` markup. The
background is black, and the `Viewport` clips text outside its bounds.

- Hardware: M5Stack Cardputer ADV
- Public API: `Viewport`, `SetText`, `SetHorizontalScroll`, `ScrollSpeed`, `ScrollGap`, `ScrollLoop`, `ScrollFromRight`, `Update`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Scroll: right-to-left loop at 24 px/s with a 32 px gap
- Expected display: the four colored language names repeat as one ticker
- Controls: none

```sh
tinygo build -target=m5stamp-s3a ./multilingual-ticker
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-ticker
```

# multilingual-showcase

This Cardputer ADV example displays nine multilingual greetings as one centered
24-dot ticker. Efont 24 renders English, Japanese, Simplified Chinese,
Traditional Chinese, Korean, Portuguese, Croatian, Russian, and Greek text with
alternating white and cyan styles. The ticker enters from the right at 50 px/s
and scrolls left once.

- Hardware: M5Stack Cardputer ADV
- Main APIs: `New`, `WithStyles`, `Viewport`, `SetText`, `SetHorizontalScroll`, `ScrollSpeed`, `ScrollFromRight`, `Update`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont24`
- Expected display: one multilingual greeting ticker centered vertically on a black background
- Controls: none

```sh
tinygo build -target=m5stamp-s3a ./multilingual-showcase
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-showcase
```

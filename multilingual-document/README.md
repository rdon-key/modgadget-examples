# multilingual-document

This example presents a longer ModGadget introduction in English, Japanese,
Chinese, and Korean. Each heading and its text use a language-specific color:
red, yellow, green, and cyan respectively. The title is white and the
background is black.

The application displays an eight-line window and advances it by one line every
1.5 seconds. It reverses direction at the beginning and end of the document,
while the `Viewport` clips the visible region.

- Hardware: M5Stack Cardputer ADV
- Main APIs: `New`, `WithStyles`, `Viewport`, `SetText`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Expected display: the colored document advances one line at a time and reverses at each end
- Controls: none; scrolling is automatic, and the example uses no horizontal scroll, audio, or keyboard input

```sh
tinygo build -target=m5stamp-s3a ./multilingual-document
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-document
```

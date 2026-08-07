# multilingual-text

最初に試す example です。英語、日本語、中国語、韓国語を一つの
`Viewport` に静的表示します。scroll、keyboard、audio は使用しません。

- Hardware: M5Stack Cardputer ADV
- Public API: `New`, `WithStyles`, `Bounds`, `Viewport`, `SetText`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Expected display: `English 日本語 中文 한국어`
- Operation: 起動後は表示したまま待機します。

```sh
tinygo build -target=m5stamp-s3a ./multilingual-text
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-text
```

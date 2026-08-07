# multilingual-document

英語、日本語、中国語、韓国語による長いModGadget紹介文を小さな画面で
読み進めるexampleです。公開APIにはvertical scroll primitiveがないため、
一つの`Viewport`に表示する8行のwindowを1.5秒ごとに上下へ1行移動します。
`Viewport`は文書を画面内へclipします。

- Hardware: M5Stack Cardputer ADV
- Public API: `New`, `WithStyles`, `Viewport`, `SetText`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Expected display: 4言語の文書が一行ずつ下方向へ進み、末尾で上方向へ戻ります。
- Operation: 自動scrollのため操作は不要です。横scroll、audio、keyboardは使用しません。

```sh
tinygo build -target=m5stamp-s3a ./multilingual-document
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-document
```

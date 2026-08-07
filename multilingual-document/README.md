# multilingual-document

英語、日本語、中国語、韓国語による長いModGadget紹介文を小さな画面で
読み進めるexampleです。Englishは赤、日本語は黄色、中文は緑、한국어は
水色で、見出しと本文を同じ色にしています。背景とタイトルは黒／白です。

現在の公開APIにはvertical offset、実行中のViewport位置変更、または公開
Surfaceがありません。`ScrollTo`もhorizontal offset専用です。そのため、
全文layoutを保持したpixel単位のsmooth vertical scrollは実装できません。
このexampleは一つの`Viewport`に表示する8行のwindowを1.5秒ごとに上下へ
1行移動する既存方式を維持し、`Viewport`が文書を画面内へclipします。

- Hardware: M5Stack Cardputer ADV
- Public API: `New`, `WithStyles`, `Viewport`, `SetText`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Expected display: 色分けした4言語の文書が一行ずつ下方向へ進み、末尾で上方向へ戻ります。
- Operation: 自動scrollのため操作は不要です。横scroll、audio、keyboardは使用しません。

```sh
tinygo build -target=m5stamp-s3a ./multilingual-document
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-document
```

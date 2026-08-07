# multilingual-ticker

`English 日本語 中文 한국어`を、4つのstyleを含む1本のloop tickerとして
画面右外から左へ流します。Englishは赤、日本語は黄色、中文は緑、한국어は
公開`<b>` markupによるbold緑です。背景はすべて黒です。`Viewport`が
画面外をclipし、ModGadgetの公開horizontal scroll APIが移動を処理します。

- Hardware: M5Stack Cardputer ADV
- Public API: `Viewport`, `SetText`, `SetHorizontalScroll`, `ScrollSpeed`, `ScrollGap`, `ScrollLoop`, `ScrollFromRight`, `Update`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Scroll: 24px/s、gap 32px、右から左へのloop
- Expected display: 色分けされた`English 日本語 中文 한국어`が1本のtickerとして繰り返し流れます。
- Operation: 操作は不要です。

```sh
tinygo build -target=m5stamp-s3a ./multilingual-ticker
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-ticker
```

# multilingual-ticker

4言語のUnicode文字列を、画面右外から左へ流れるloop tickerとして表示します。
`Viewport` が画面外をclipし、ModGadgetの公開horizontal scroll APIが移動を
処理します。

- Hardware: M5Stack Cardputer ADV
- Public API: `Viewport`, `SetText`, `SetHorizontalScroll`, `ScrollSpeed`, `ScrollGap`, `ScrollLoop`, `ScrollFromRight`, `Update`, `Render`
- Font: `github.com/rdon-key/modgadget-fonts/efont16`
- Expected display: `English 日本語 中文 한국어` が繰り返し右から左へ流れます。
- Operation: 操作は不要です。

```sh
tinygo build -target=m5stamp-s3a ./multilingual-ticker
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-ticker
```

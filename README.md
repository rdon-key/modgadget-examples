# ModGadget Examples

[ModGadget](https://github.com/rdon-key/modgadget)を使うCardputer ADV向けの
TinyGo example集です。ModGadgetと`modgadget-fonts`は通常のGo module依存
として取得されるため、ModGadget本体のcloneは不要です。各exampleは独立して
buildでき、静的表示からkeyboard applicationまで順番に試せます。

## Examples

| Example | 内容 |
| --- | --- | --- |
| [`multilingual-text`](multilingual-text/) | 英語・日本語・中国語・韓国語の静的表示 |
| [`multilingual-ticker`](multilingual-ticker/) | 4言語文字列のloop horizontal scroll |
| [`multilingual-ticker2`](multilingual-ticker2/) | 24dotの4言語tickerを別々の位置・速度で表示し、elapsed timerを表示 |
| [`multilingual-document`](multilingual-document/) | 4言語のModGadget紹介文をvertical scroll |
| [`audio-beep`](audio-beep/) | Cardputer ADVで音を再生 |
| [`keyboard-typing`](keyboard-typing/) | keyboardを使ったtyping application |

## Test

Host側で確認できるlogic testは次で実行します。hardware用`main.go`はTinyGo
build tag付きです。

```sh
go test ./...
```

## TinyGo build

Cardputer ADV対応の`m5stamp-s3a` targetを含むTinyGo開発版が必要です。
例えば最初のexampleは次のようにbuildします。

```sh
tinygo build -target=m5stamp-s3a ./multilingual-text
```

他のexampleも末尾のdirectory名を置き換えて単独でbuildできます。flashは
各exampleのREADMEに記載した別操作です。

## License

BSD 3-Clause.

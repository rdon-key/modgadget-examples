# audio-beep

Cardputer ADVの公開audio APIで、組み込みsound patternを一度ずつ再生する
最小exampleです。cooperative playerを`Update`しながら各音の完了を待ち、
sequence終了後は静かに待機します。

- Hardware: M5Stack Cardputer ADV
- Public API: `device/cardputeradv.ConfigureAudio`, `audio/cardputeradv.Player.PlayPattern`, `Update`, `Busy`, `Stop`
- Font: 使用しません。
- Expected sound: Startup、Click、Correctを各1回、300ms間隔で再生します。
- Operation: 起動後は操作不要です。displayとkeyboardは使用しません。

```sh
tinygo build -target=m5stamp-s3a ./audio-beep
tinygo flash -target=m5stamp-s3a -port=<PORT> ./audio-beep
```

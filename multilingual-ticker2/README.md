# multilingual-ticker2

Cardputer ADVの1画面を複数領域として使い、Efont 24でEnglish、日本語、中文、
한국어の4つの独立tickerを異なる位置・速度・gap・色で右から左へ流します。
各`Viewport`がそれぞれの範囲で表示をclipします。最下段には白いEfont 16で
起動後の経過時間を`MM:SS`表示します。

| Text | Bounds `(x, y, width, height)` | Color | Speed | Gap |
| --- | --- | --- | --- | --- |
| English | `(0, 0, 240, 24)` | red | 20px/s | 40px |
| 日本語 | `(24, 27, 216, 24)` | yellow | 24px/s | 48px |
| 中文 | `(0, 54, 190, 24)` | green | 28px/s | 56px |
| 한국어 | `(40, 81, 200, 24)` | cyan | 32px/s | 64px |

期待される画面は、独立してloopする4本の24dot tickerと、右下で1秒ごとに
更新される`00:00`形式のelapsed timerです。操作は不要です。

```sh
tinygo build -target=m5stamp-s3a ./multilingual-ticker2
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-ticker2
```

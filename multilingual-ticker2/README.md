# multilingual-ticker2

This Cardputer ADV example divides one screen into multiple regions. Four
independent Efont 24 tickers scroll English, Japanese, Chinese, and Korean from
right to left with different speeds, gaps, and colors. Each `Viewport` clips
its own text. A white Efont 16 elapsed timer appears at the bottom.

| Text | Bounds `(x, y, width, height)` | Color | Speed | Gap |
| --- | --- | --- | --- | --- |
| English | `(0, 0, 240, 24)` | red | 20px/s | 40px |
| 日本語 | `(0, 27, 240, 24)` | yellow | 24px/s | 48px |
| 中文 | `(0, 54, 240, 24)` | green | 60px/s | 56px |
| 한국어 | `(0, 81, 240, 24)` | cyan | 50px/s | 64px |

The expected display contains four independently looping 24-dot tickers and a
right-aligned `MM:SS` timer updated once per second. No controls are required.

```sh
tinygo build -target=m5stamp-s3a ./multilingual-ticker2
tinygo flash -target=m5stamp-s3a -port=<PORT> ./multilingual-ticker2
```

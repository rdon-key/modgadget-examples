//go:build tinygo

package main

import (
	"strings"
	"time"

	"github.com/rdon-key/modgadget"
	"github.com/rdon-key/modgadget-fonts/efont16"
	board "github.com/rdon-key/modgadget/device/cardputeradv"
)

const visibleLines = 8

var documentLines = [...]string{
	"ModGadget",
	"",
	"English",
	"ModGadget is a UI foundation",
	"for TinyGo displays.",
	"It helps build structured UIs",
	"on small microcontrollers.",
	"Viewports divide the display",
	"into clipped regions.",
	"Unicode fonts show multiple",
	"languages and text can scroll.",
	"It runs on Cardputer ADV and",
	"compact devices. Goal:",
	"application-level UI on limited",
	"hardware.",
	"",
	"日本語",
	"ModGadgetはTinyGo向けの",
	"UIと表示の基盤です。小型の",
	"マイコンでも構造化した画面を",
	"組み立てやすくします。Viewport",
	"を使うと、表示を独立した領域に",
	"分割し、それぞれを範囲内に収め",
	"られます。Unicode対応の",
	"フォントで多言語を表示でき、",
	"文章のスクロールも扱えます。",
	"Cardputer ADVなどの小型",
	"デバイスで、限られた資源でも",
	"アプリケーションらしいUIを作る",
	"ことを目的としています。",
	"",
	"中文",
	"ModGadget是面向TinyGo的UI",
	"与显示基础。它让小型微控制器",
	"更容易构建结构化界面。使用",
	"Viewport可以把屏幕分成多个",
	"独立区域，并把绘制内容裁剪在",
	"区域之内。它能使用Unicode",
	"多语言字体，也能处理文字滚动。",
	"它可用于Cardputer ADV等",
	"小型设备，目标是在资源有限的",
	"在资源有限的硬件上也能构建具有",
	"应用程序层次的实用界面。",
	"",
	"한국어",
	"ModGadget은 TinyGo용 UI와",
	"디스플레이 기반입니다. 작은",
	"마이크로컨트롤러에서도",
	"구조화된",
	"화면을 쉽게 만들도록 돕습니다.",
	"Viewport로 화면을 독립",
	"영역으로 나누고 그 범위 안에서",
	"그립니다.",
	"Unicode 다국어 글꼴과 텍스트",
	"스크롤을 사용할 수 있습니다.",
	"Cardputer ADV 같은 소형",
	"장치에서 제한된 자원으로도",
	"응용 프로그램 수준의 UI를",
	"만드는 것이 목표입니다.",
}

func visibleDocument(first int) string {
	var text strings.Builder
	for index := 0; index < visibleLines; index++ {
		if index > 0 {
			text.WriteString("<br>")
		}
		text.WriteString(documentLines[first+index])
	}
	return text.String()
}

func main() {
	time.Sleep(3 * time.Second)
	panel, err := board.ConfigureDisplay()
	if err != nil {
		panic(err)
	}
	gadget := modgadget.New(panel, modgadget.WithStyles(modgadget.StyleSet{
		Default: modgadget.Style{Font: efont16.Font, Foreground: modgadget.ColorWhite, Background: modgadget.ColorBlack},
	}))
	if err := gadget.Clear(); err != nil {
		panic(err)
	}
	view := gadget.Viewport(modgadget.Bounds(0, 0, board.DisplayWidth, board.DisplayHeight))

	first, direction := 0, 1
	for {
		if err := view.SetText(visibleDocument(first)); err != nil {
			panic(err)
		}
		if err := gadget.Render(); err != nil {
			panic(err)
		}
		time.Sleep(1500 * time.Millisecond)

		first += direction
		if first == len(documentLines)-visibleLines || first == 0 {
			direction = -direction
		}
	}
}

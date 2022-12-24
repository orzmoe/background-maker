package background_maker

import (
	"background-maker/pkg/trianglify"
	"github.com/fogleman/gg"
	"testing"
)

func TestBackgroundMaker(t *testing.T) {
	fontFace, err := gg.LoadFontFace("PingFang SC Bold.ttf", 42)
	if err != nil {
		return
	}

	dc, err := trianglify.New(800, 450, 3000, "癸卯年9月1日七月十七出生男孩什么命缺火宝宝名字", fontFace, nil)
	if err != nil {
		t.Log(err)
	}
	dc.SavePNG("triangle.png")
}

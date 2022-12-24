package background_maker

import (
	"background-maker/pkg/trianglify"
	"github.com/fogleman/gg"
)

func BackgroundMaker(text string) (*gg.Context, error) {
	fontFace, err := gg.LoadFontFace("../PingFang SC Bold.ttf", 48)
	if err != nil {
		return nil, err
	}
	return trianglify.New(800, 450, 3000, text, fontFace, nil)
}

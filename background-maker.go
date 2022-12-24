package background_maker

import (
	"github.com/fogleman/gg"
	"github.com/orzmoe/background-maker/pkg/trianglify"
)

func BackgroundMaker(text string) (*gg.Context, error) {
	fontFace, err := gg.LoadFontFace("../PingFang SC Bold.ttf", 48)
	if err != nil {
		return nil, err
	}
	return trianglify.New(800, 450, 3000, text, fontFace, nil)
}

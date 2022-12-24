package background_maker

import (
	_ "embed"
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/orzmoe/background-maker/pkg/trianglify"
)

//go:embed "PingFang SC Bold.ttf"
var fontEmbed []byte

func BackgroundMaker(text string) (*gg.Context, error) {
	font, err := freetype.ParseFont(fontEmbed)
	if err != nil {
		return nil, err
	}
	fontFace := truetype.NewFace(font, &truetype.Options{
		Size: 42,
	})
	return trianglify.New(800, 450, 3000, text, fontFace, nil)
}

package trianglify

import (
	"image/color"
	"strconv"
)

type Colors struct {
	// 背景色
	Background []color.Color `json:"background"`
	// 字体颜色
	Foreground color.Color `json:"foreground"`
}

var DefaultColors = []Colors{
	{
		Background: []color.Color{
			color.RGBA{R: 199, G: 82, B: 42, A: 255},
			color.RGBA{R: 229, G: 193, B: 133, A: 255},
			color.RGBA{R: 251, G: 242, B: 196, A: 255},
			color.RGBA{R: 116, G: 168, B: 146, A: 255},
			color.RGBA{R: 0, G: 133, B: 133, A: 255},
		},
		Foreground: color.RGBA{R: 90, G: 32, B: 7, A: 255},
	},
	{
		Background: []color.Color{
			color.RGBA{R: 249, G: 224, B: 192, A: 255},
			color.RGBA{R: 24, G: 69, B: 117, A: 255},
		},
		Foreground: HexToRGB("#c17412"),
	},
	{
		Background: []color.Color{
			HexToRGB("809bce"),
			HexToRGB("95b8d1"),
			HexToRGB("b8e0d4"),
			HexToRGB("d6eadf"),
			HexToRGB("eac4d5"),
		},
		Foreground: HexToRGB("#367c67"),
	},
	{
		Background: []color.Color{
			HexToRGB("c095e4"),
			HexToRGB("fcedf2"),
			HexToRGB("ffd1d4"),
			HexToRGB("ffb7c5"),
			HexToRGB("ffa0c5"),
		},
		Foreground: HexToRGB("#d00029"),
	},
	{
		Background: []color.Color{
			HexToRGB("cdb4db"),
			HexToRGB("ffc8dd"),
			HexToRGB("faaac7"),
			HexToRGB("bee2ff"),
			HexToRGB("a2d2ff"),
		},
		Foreground: HexToRGB("#d00029"),
	},
}

// HexToRGB 16进制颜色转换为RGB
func HexToRGB(hex string) color.Color {
	var r, g, b uint8
	if hex[0] == '#' {
		hex = hex[1:]
	}
	rgb, _ := strconv.ParseUint(hex, 16, 64)
	r = uint8(rgb >> 16)
	g = uint8(rgb >> 8)
	b = uint8(rgb)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

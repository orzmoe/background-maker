package trianglify

import (
	"github.com/fogleman/delaunay"
	"github.com/fogleman/gg"
	"github.com/fogleman/poissondisc"
	"golang.org/x/image/font"
	"image/color"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode"
)

func New(width, height, n int, text string, fontFace font.Face, foreground color.Color, colors ...color.Color) (*gg.Context, error) {
	if len(colors) == 0 {
		//DefaultColors 随机取一个
		defaultColor := DefaultColors[RandomInt(0, len(DefaultColors))]
		colors = defaultColor.Background
		if foreground == nil {
			foreground = defaultColor.Foreground
		}
	}
	if foreground == nil {
		foreground = color.Black
	}
	dc := NewGradient(width, height, colors...)

	// 复制一份dc
	dc2 := NewGradient(width, height, colors...)
	// generate points
	// 创建一个随机点集
	points := generatePoints(width, height, n)
	// triangulate
	// 三角化
	triangulation, err := delaunay.Triangulate(points)

	if err != nil {
		return nil, err
	}
	// compute point bounds for rendering
	// 计算点的边界
	min := points[0]
	max := points[0]
	for _, p := range points {
		min.X = math.Min(min.X, p.X)
		min.Y = math.Min(min.Y, p.Y)
		max.X = math.Max(max.X, p.X)
		max.Y = math.Max(max.Y, p.Y)
	}

	size := delaunay.Point{X: max.X - min.X, Y: max.Y - min.Y}
	center := delaunay.Point{X: min.X + size.X/2, Y: min.Y + size.Y/2}
	scale := math.Min(float64(width)/size.X, float64(height)/size.Y) * 1.15

	dc.Translate(float64(width/2), float64(height/2))
	dc.Scale(scale, scale)
	dc.Translate(-center.X, -center.Y)

	ts := triangulation.Triangles
	hs := triangulation.Halfedges
	for i, h := range hs {
		if i > h {
			// 获取三角形的三个顶点
			p1 := points[ts[i]]
			p2 := points[ts[nextHalfEdge(i)]]
			p3 := points[ts[nextHalfEdge(nextHalfEdge(i))]]
			//绘制三角形
			dc.MoveTo(p1.X, p1.Y)
			dc.LineTo(p2.X, p2.Y)
			dc.LineTo(p3.X, p3.Y)
			dc.ClosePath()
			// 设置填充样式为p1.X, p1.Y点的颜色
			dc.SetFillStyle(gg.NewSolidPattern(getTriangleColor(dc2, p1, p2, p3)))
			dc.Fill()
		}
	}
	dc.SetFontFace(fontFace)
	dc.SetColor(foreground)
	//去除空格
	text = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, text)
	var texts []string
	line := 0
	count := 0
	for _, x := range text {
		if count >= 21 { // 一行最大显示的字符数量
			line++
			count = 0
		}
		if len(texts) <= line {
			texts = append(texts, string(x))
		} else {
			texts[line] = texts[line] + string(x)
		}

		//汉字+2，字母+1
		if unicode.Is(unicode.Han, x) {
			count += 2
		} else {
			count++
		}
	}
	if len(texts) > 3 {
		texts = texts[:3]
		texts = append(texts, "...")
	}
	newText := ""
	for i, t := range texts {
		newText += t
		if i != len(texts)-1 {
			newText += " "
		}
	}
	dc.DrawStringWrapped(newText, float64(width/2), float64(height/2), 0.5, 0.5, float64(width/2), 1.2, gg.AlignCenter)
	return dc, nil
}

func RandomInt(i int, i2 int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i2-i) + i
}

// 获取三角形内平均颜色
func getTriangleColor(dc *gg.Context, p1, p2, p3 delaunay.Point) color.RGBA {
	// 计算三角形的三个顶点的中点
	p1x := (p1.X + p2.X + p3.X) / 3
	p1y := (p1.Y + p2.Y + p3.Y) / 3

	// 获取p1x, p1y点的颜色
	r, g, b, _ := dc.Image().At(int(p1x), int(p1y)).RGBA()
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func NewGradient(width, height int, colors ...color.Color) *gg.Context {
	dc := gg.NewContext(width, height)
	grad := gg.NewLinearGradient(0, 0, float64(width), float64(height))
	for i, c := range colors {
		if i == 0 {
			grad.AddColorStop(0, c)
		}
		if i == len(colors)-1 {
			grad.AddColorStop(1, c)
		} else {
			grad.AddColorStop(float64(i)/float64(len(colors)-1), c)
		}
	}
	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	return dc
}
func generatePoints(w, h, n int) []delaunay.Point {
	points := poissondisc.Sample(0, 0, float64(w), float64(h), 46, 12, nil)
	sort.Slice(points, func(i, j int) bool {
		p1 := points[i]
		p2 := points[j]
		d1 := math.Hypot(p1.X, p1.Y)
		d2 := math.Hypot(p2.X, p2.Y)
		return d1 < d2
	})
	if len(points) > n {
		points = points[:n]
	}
	result := make([]delaunay.Point, len(points))
	for i, p := range points {
		result[i].X = p.X
		result[i].Y = p.Y
	}
	return result
}
func nextHalfEdge(e int) int {
	if e%3 == 2 {
		return e - 2
	}
	return e + 1
}

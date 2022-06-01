package chart

import (
	"github.com/rarnu/golcl/lcl"
	"github.com/rarnu/golcl/lcl/bitmap"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
	"github.com/vdobler/chart/txtg"
	"image"
	"image/color"
	"image/draw"
)

type Dumper struct {
	N, M, W, H, Cnt int
	I               *image.RGBA
	image           *lcl.TImage
	changeSize      bool
}

func NewDumper(col, row, w, h int, img *lcl.TImage, changeSize bool) *Dumper {
	dumper := Dumper{N: col, M: row, W: w, H: h}
	dumper.image = img
	dumper.changeSize = changeSize
	dumper.I = image.NewRGBA(image.Rect(0, 0, col*w, row*h))
	bg := image.NewUniform(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.ZP, draw.Src)

	return &dumper
}
func (d *Dumper) Close() {
	if d.image != nil {
		if d.changeSize {
			d.image.Picture().Clear()
			// 必须先改大小才能画，否则画出框
			w := int32(d.N * d.W)
			h := int32(d.M * d.H)
			d.image.SetWidth(w)
			d.image.SetHeight(h)
		}
		pngObj, err := bitmap.ToPngImage(d.I)
		if err == nil && pngObj != nil {
			defer pngObj.Free()
			d.image.Canvas().Draw(0, 0, pngObj)
		}
	}
}

func (d *Dumper) Plot(c chart.Chart) {
	row, col := d.Cnt/d.N, d.Cnt%d.N
	igr := imgg.AddTo(d.I, col*d.W, row*d.H, d.W, d.H, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, nil, nil)
	c.Plot(igr)
	tgr := txtg.New(100, 30)
	c.Plot(tgr)
	d.Cnt++
}

package widgets

import (
	. "github.com/kallaurru/termui/v3"
	"image"
)

type GaugeFloat struct {
	Block
	gauge       *Gauge
	widgetTheme *WidgetTheme
}

func NewGaugeFloat(center image.Point, title string) *GaugeFloat {
	gf := &GaugeFloat{
		Block: *NewBlock(),
		gauge: NewGauge(),
	}

	return gf
}

func (gf *GaugeFloat) SetGaugeTheme(theme *GaugeTheme) {
	gf.gauge.UploadTheme(theme)
}

func (gf *GaugeFloat) Draw(buf *Buffer) {
	gf.Block.Draw(buf)
	gf.gauge.Draw(buf)
}

func (gf *GaugeFloat) makeBlockRect(center image.Point) {
	/*	half := limitMsgLine / 2
		xMin := center.X - half - 2
		xMax := center.X + half + 2
		yMax := db.button.Block.Max.Y + 2
		yMin := center.Y - 8
		db.Block.SetRect(xMin, yMin, xMax, yMax)
		db.msg.Block.SetRect(xMin+1, yMin+2, xMax-1, center.Y-2)
	*/
}

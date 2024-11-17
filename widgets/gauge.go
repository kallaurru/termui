// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"fmt"
	"image"

	. "github.com/kallaurru/termui/v3"
)

type Gauge struct {
	Block
	Percent    int
	BarColor   Color
	Label      string
	LabelStyle Style
}

func NewGauge() *Gauge {
	return &Gauge{
		Block:      *NewBlock(),
		BarColor:   Theme.Gauge.Bar,
		LabelStyle: Theme.Gauge.Label,
	}
}

func (g *Gauge) UploadTheme(gt *GaugeTheme) {
	g.LabelStyle = gt.Label
	g.BarColor = gt.Label.Fg
}

func (g *Gauge) UploadWidgetTheme(wt *WidgetTheme, isActive bool) {
	g.Block.BorderStyle = wt.GetBorderStyle(isActive)
	g.Block.TitleStyle = wt.GetTitleStyle(isActive)
}

func (g *Gauge) Draw(buf *Buffer) {
	g.Block.Draw(buf)

	label := g.Label
	if label == "" {
		label = fmt.Sprintf("%d%%", g.Percent)
	}

	// plot bar
	barWidth := int((float64(g.Percent) / 100) * float64(g.Inner.Dx()))
	buf.Fill(
		NewCell(' ', NewStyle(ColorClear, g.BarColor)),
		image.Rect(g.Inner.Min.X, g.Inner.Min.Y, g.Inner.Min.X+barWidth, g.Inner.Max.Y),
	)

	// plot label
	labelXCoordinate := g.Inner.Min.X + (g.Inner.Dx() / 2) - int(float64(len(label))/2)
	labelYCoordinate := g.Inner.Min.Y + ((g.Inner.Dy() - 1) / 2)
	if labelYCoordinate < g.Inner.Max.Y {
		for i, char := range label {
			style := g.LabelStyle
			if labelXCoordinate+i+1 <= g.Inner.Min.X+barWidth {
				style = NewStyle(g.BarColor, ColorClear, ModifierReverse)
			}
			buf.SetCell(NewCell(char, style), image.Pt(labelXCoordinate+i, labelYCoordinate))
		}
	}
}

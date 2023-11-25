// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

var StandardColors = []Color{
	ColorRed,
	ColorGreen,
	ColorYellow,
	ColorBlue,
	ColorMagenta,
	ColorCyan,
	ColorWhite,
}

var StandardStyles = []Style{
	NewStyle(ColorRed),
	NewStyle(ColorGreen),
	NewStyle(ColorYellow),
	NewStyle(ColorBlue),
	NewStyle(ColorMagenta),
	NewStyle(ColorCyan),
	NewStyle(ColorWhite),
}

type WidgetTheme struct {
	titleStyle        Style
	borderStyle       Style
	activeTitleStyle  Style
	activeBorderStyle Style
}

type RootTheme struct {
	Default Style

	Block BlockTheme

	BarChart        BarChartTheme
	Gauge           GaugeTheme
	Plot            PlotTheme
	List            ListTheme
	Tree            TreeTheme
	Paragraph       ParagraphTheme
	PieChart        PieChartTheme
	Sparkline       SparklineTheme
	StackedBarChart StackedBarChartTheme
	Tab             TabTheme
	Table           TableTheme
}

type BlockTheme struct {
	Title  Style
	Border Style
}

type BarChartTheme struct {
	Bars   []Color
	Nums   []Style
	Labels []Style
}

type GaugeTheme struct {
	Bar   Color
	Label Style
}

type PlotTheme struct {
	Lines []Color
	Axes  Color
}

type ListTheme struct {
	Text Style
}

type TreeTheme struct {
	Text      Style
	Collapsed rune
	Expanded  rune
}

type ParagraphTheme struct {
	Text Style
}

type PieChartTheme struct {
	Slices []Color
}

type SparklineTheme struct {
	Title Style
	Line  Color
}

type StackedBarChartTheme struct {
	Bars   []Color
	Nums   []Style
	Labels []Style
}

type TabTheme struct {
	Active   Style
	Inactive Style
}

type TableTheme struct {
	Text Style
}

// Theme holds the default Styles and Colors for all widgets.
// You can set default widget Styles by modifying the Theme before creating the widgets.
var Theme = RootTheme{
	Default: NewStyle(ColorWhite),

	Block: BlockTheme{
		Title:  NewStyle(ColorWhite),
		Border: NewStyle(ColorWhite),
	},

	BarChart: BarChartTheme{
		Bars:   StandardColors,
		Nums:   StandardStyles,
		Labels: StandardStyles,
	},

	Paragraph: ParagraphTheme{
		Text: NewStyle(ColorWhite),
	},

	PieChart: PieChartTheme{
		Slices: StandardColors,
	},

	List: ListTheme{
		Text: NewStyle(ColorWhite),
	},

	Tree: TreeTheme{
		Text:      NewStyle(ColorWhite),
		Collapsed: COLLAPSED,
		Expanded:  EXPANDED,
	},

	StackedBarChart: StackedBarChartTheme{
		Bars:   StandardColors,
		Nums:   StandardStyles,
		Labels: StandardStyles,
	},

	Gauge: GaugeTheme{
		Bar:   ColorWhite,
		Label: NewStyle(ColorWhite),
	},

	Sparkline: SparklineTheme{
		Title: NewStyle(ColorWhite),
		Line:  ColorWhite,
	},

	Plot: PlotTheme{
		Lines: StandardColors,
		Axes:  ColorWhite,
	},

	Table: TableTheme{
		Text: NewStyle(ColorWhite),
	},

	Tab: TabTheme{
		Active:   NewStyle(ColorRed),
		Inactive: NewStyle(ColorWhite),
	},
}

// GetTitleStyle - получить стиль заголовка
func (wt *WidgetTheme) GetTitleStyle(isActive bool) Style {
	if isActive {
		return wt.activeTitleStyle
	}

	return wt.titleStyle
}

// GetBorderStyle - получить стиль линии обводки
func (wt *WidgetTheme) GetBorderStyle(isActive bool) Style {
	if isActive {
		return wt.activeBorderStyle
	}

	return wt.borderStyle
}

func NewDefWidgetTheme() *WidgetTheme {
	return &WidgetTheme{
		titleStyle:        NewStyle(ColorBlack, ColorWhite),
		borderStyle:       NewStyle(ColorWhite),
		activeTitleStyle:  NewStyle(ColorBlack, ColorWhite, ModifierBold),
		activeBorderStyle: NewStyle(ColorGreenBlue),
	}
}

func NewMyDefaultWidgetTheme() *WidgetTheme {
	return &WidgetTheme{
		titleStyle:        NewStyle(ColorWhite),
		borderStyle:       NewStyle(ColorWhite),
		activeTitleStyle:  NewStyle(ColorGreenBlue, ColorClear, ModifierBold),
		activeBorderStyle: NewStyle(ColorGreenBlue),
	}
}

func NewButtonDefaultTheme() *WidgetTheme {
	return &WidgetTheme{
		titleStyle:        NewStyle(ColorWhite),
		borderStyle:       NewStyle(ColorWhite),
		activeTitleStyle:  NewStyleBgFree(ColorGreenBlue, ModifierBold),
		activeBorderStyle: NewStyle(ColorGreenBlue),
	}
}

package termui

// Color is an integer from -1 to 255
// -1 = ColorClear
// 0-255 = Xterm colors
type Color int

// ColorClear clears the Fg or Bg color of a Style
const ColorClear Color = -1

// Basic terminal colors
const (
	ColorBlack         Color = 0
	ColorRed           Color = 1
	ColorGreen         Color = 2
	ColorYellow        Color = 3
	ColorBlue          Color = 4
	ColorMagenta       Color = 5
	ColorCyan          Color = 6
	ColorWhite         Color = 7
	ColorGreenBgDunkel Color = 22
	ColorGreenBlue     Color = 23
	ColorBlueLightFg   Color = 30
	ColorBlueLightFg2  Color = 31
)

type Modifier uint

const (
	// ModifierClear clears any modifiers
	ModifierClear     Modifier = 0
	ModifierBold      Modifier = 1 << 9
	ModifierUnderline Modifier = 1 << 10
	ModifierReverse   Modifier = 1 << 11
)

// Style represents the style of one terminal cell
type Style struct {
	Fg       Color
	Bg       Color
	Modifier Modifier
}

// StyleClear represents a default Style, with no colors or modifiers
var StyleClear = Style{
	Fg:       ColorClear,
	Bg:       ColorClear,
	Modifier: ModifierClear,
}

// NewStyle takes 1 to 3 arguments
// 1st argument = Fg
// 2nd argument = optional Bg
// 3rd argument = optional Modifier
func NewStyle(fg Color, args ...interface{}) Style {
	st := Style{
		Fg:       fg,
		Bg:       ColorClear,
		Modifier: ModifierClear,
	}
	if len(args) >= 1 {
		_, isColor := args[0].(Color)
		if isColor {
			st.Bg = args[0].(Color)
		} else {
			// что бы можно было добавить модификатор не меняя родного bg терминала
			st.Modifier = args[0].(Modifier)
		}

		return st
	}
	if len(args) == 2 {
		st.Modifier = args[1].(Modifier)
	}

	return st
}

func NewStyleBgFree(fg Color, mod Modifier) Style {
	return NewStyle(fg, mod)
}

func NewStyleFromStyleCode(code uint32) Style {
	// 1 << 24 fg = -1
	// 1 << 25 bg = -1
	// младший байт = fg
	// второй байт = modifier
	// третий байт = bg
	var bg, fg Color
	var mod Modifier
	// сначала проверим есть ли значения -1
	negative := code >> 24
	mod = Modifier(code & 0x0000ff00)

	switch negative {
	case 0:
		fg = Color(code & 0xff)
		bg = Color((code >> 16) & 0xff)
	case 1:
		fg = -1 * Color(negative)
		bg = Color((code >> 16) & 0xff)
	case 2:
		fg = -1 * Color(negative)
		bg = -1 * Color(negative)
	}

	return Style{
		Bg:       bg,
		Fg:       fg,
		Modifier: mod,
	}
}

func StyleSerializeToCode(st Style) uint32 {
	var code uint32

	if st.Fg == -1 {
		code |= 1 << 24
	} else {
		code |= uint32(st.Fg)
	}
	if st.Bg == -1 {
		code |= 1 << 25
	} else {
		code |= uint32(st.Bg) << 16
	}

	code |= uint32(st.Modifier)
	return code
}

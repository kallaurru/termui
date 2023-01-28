// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"fmt"
	"strings"
)

const (
	tokenFg       = "fg"
	tokenBg       = "bg"
	tokenModifier = "mod"

	tokenItemSeparator  = ","
	tokenValueSeparator = ":"

	TokenBeginStyledText = '['
	TokenEndStyledText   = ']'

	tokenBeginStyle = '('
	tokenEndStyle   = ')'
)

type parserState uint

const (
	parserStateDefault parserState = iota
	parserStateStyleItems
	parserStateStyledText
)

// StyleParserColorMap can be modified to add custom color parsing to text
var StyleParserColorMap = map[string]Color{
	"red":     ColorRed,
	"blue":    ColorBlue,
	"black":   ColorBlack,
	"cyan":    ColorCyan,
	"yellow":  ColorYellow,
	"white":   ColorWhite,
	"clear":   ColorClear,
	"green":   ColorGreen,
	"magenta": ColorMagenta,
	"22":      ColorGreenBgDunkel,
	"23":      ColorGreenBlue,
	"30":      ColorBlueLightFg,
	"31":      ColorBlueLightFg2,
}

var modifierMap = map[string]Modifier{
	"bold":      ModifierBold,
	"underline": ModifierUnderline,
	"reverse":   ModifierReverse,
}

func StyleToString(style Style) string {
	var (
		hasFg, hasBg, hasMod = false, false, false
		out                  = ""
	)

	m := make(map[Color]string)
	mm := make(map[Modifier]string)

	for strColor, color := range StyleParserColorMap {
		m[color] = strColor
	}

	for strMod, mod := range modifierMap {
		mm[mod] = strMod
	}

	fg, ok := m[style.Fg]
	if ok {
		hasFg = true
	}

	bg, ok := m[style.Bg]
	if ok {
		hasBg = true
	}

	mod, ok := mm[style.Modifier]
	if ok {
		hasMod = true
	}
	if hasFg {
		out = fmt.Sprintf("%s%s%s", tokenFg, tokenValueSeparator, fg)
	}

	if hasBg {
		if len(out) > 0 {
			out = fmt.Sprintf("%s%s%s%s%s", out, tokenItemSeparator, tokenBg, tokenValueSeparator, bg)
		} else {
			out = fmt.Sprintf("%s%s%s", tokenBg, tokenValueSeparator, bg)
		}
	}

	if hasMod {
		if len(out) > 0 {
			out = fmt.Sprintf("%s%s%s%s%s", out, tokenItemSeparator, tokenBg, tokenValueSeparator, mod)
		} else {
			out = fmt.Sprintf("%s%s%s", tokenModifier, tokenValueSeparator, mod)
		}
	}

	if len(out) > 0 {
		return fmt.Sprintf("(%s)", out)
	}

	return out
}

// readStyle translates an []rune like `fg:red,mod:bold,bg:white` to a style
func readStyle(runes []rune, defaultStyle Style) Style {
	style := defaultStyle
	split := strings.Split(string(runes), tokenItemSeparator)
	for _, item := range split {
		pair := strings.Split(item, tokenValueSeparator)
		if len(pair) == 2 {
			switch pair[0] {
			case tokenFg:
				style.Fg = StyleParserColorMap[pair[1]]
			case tokenBg:
				style.Bg = StyleParserColorMap[pair[1]]
			case tokenModifier:
				style.Modifier = modifierMap[pair[1]]
			}
		}
	}
	return style
}

// ParseStyles parses a string for embedded Styles and returns []Cell with the correct styling.
// Uses defaultStyle for any text without an embedded style.
// Syntax is of the form [text](fg:<color>,mod:<attribute>,bg:<color>).
// Ordering does not matter. All fields are optional.
func ParseStyles(s string, defaultStyle Style) []Cell {
	var cells []Cell
	var styledText []rune
	var styleItems []rune

	runes := []rune(s)
	state := parserStateDefault
	squareCount := 0

	reset := func() {
		styledText = []rune{}
		styleItems = []rune{}
		state = parserStateDefault
		squareCount = 0
	}

	rollback := func() {
		cells = append(cells, RunesToStyledCells(styledText, defaultStyle)...)
		cells = append(cells, RunesToStyledCells(styleItems, defaultStyle)...)
		reset()
	}

	// chop first and last runes
	chop := func(s []rune) []rune {
		return s[1 : len(s)-1]
	}

	for i, _rune := range runes {
		switch state {
		case parserStateDefault:
			if _rune == TokenBeginStyledText {
				state = parserStateStyledText
				squareCount = 1
				styledText = append(styledText, _rune)
			} else {
				cells = append(cells, Cell{_rune, defaultStyle})
			}
		case parserStateStyledText:
			switch {
			case squareCount == 0:
				switch _rune {
				case tokenBeginStyle:
					state = parserStateStyleItems
					styleItems = append(styleItems, _rune)
				default:
					rollback()
					switch _rune {
					case TokenBeginStyledText:
						state = parserStateStyledText
						squareCount = 1
						styleItems = append(styleItems, _rune)
					default:
						cells = append(cells, Cell{_rune, defaultStyle})
					}
				}
			case len(runes) == i+1:
				rollback()
				styledText = append(styledText, _rune)
			case _rune == TokenBeginStyledText:
				squareCount++
				styledText = append(styledText, _rune)
			case _rune == TokenEndStyledText:
				squareCount--
				styledText = append(styledText, _rune)
			default:
				styledText = append(styledText, _rune)
			}
		case parserStateStyleItems:
			styleItems = append(styleItems, _rune)
			if _rune == tokenEndStyle {
				style := readStyle(chop(styleItems), defaultStyle)
				cells = append(cells, RunesToStyledCells(chop(styledText), style)...)
				reset()
			} else if len(runes) == i+1 {
				rollback()
			}
		}
	}

	return cells
}

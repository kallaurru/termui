package widgets

import (
	. "github.com/kallaurru/termui/v3"
)

type Widget interface {
	Drawable
	GetId() string
}

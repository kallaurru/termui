package widgets

import (
	. "github.com/kallaur/termui/v3"
)

type Widget interface {
	Drawable
	GetId() string
}

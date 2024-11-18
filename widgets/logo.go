package widgets

import . "github.com/kallaurru/termui/v3"

type Logo struct {
	Block
	filename string // файл с баннером созданном в calligraphy
	bannerX  int    // длина баннера
	bannerY  int    // высота баннера
}

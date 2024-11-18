package widgets

import (
	. "github.com/kallaurru/termui/v3"
)

type Banner struct {
	Block
	widgetTheme WidgetTheme
	data        []byte // заливаем как массив байт. На строки разбираем по \n
	len         int    // длина баннера
	high        int    // ширина баннера
}

func NewBanner() *Banner {
	const size = 512 // должно перекрыть размеры большинства случаем

	banner := &Banner{
		Block:       *NewBlock(),
		data:        make([]byte, 0, size),
		len:         0,
		high:        0,
		widgetTheme: WidgetTheme{},
	}

	return banner
}

func (b *Banner) UploadFromFile(filename string) error {
	return nil
}

func (b *Banner) UploadWidgetTheme(wt WidgetTheme) {
	b.Block.BorderStyle = wt.GetBorderStyle(false)
	b.Block.TitleStyle = wt.GetTitleStyle(false)
}

func (b *Banner) Draw(buf *Buffer) {
	b.Block.Draw(buf)
}

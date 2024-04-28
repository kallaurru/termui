package tmpl

type GridSchemaBuilder struct {
	countColsPerRow   []uint8
	adaptiveColsSizes uint8
	adaptiveRowSizes  uint8
	widgets           map[uint16]string
}

// позволяет описать схему расположения виджетов строками
// добавляем строку, колонку и тип виджета который там будет

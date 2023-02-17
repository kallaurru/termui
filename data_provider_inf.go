package termui

type UpdatingDataProvider interface {
	UpdateData(data string, address ...uint32)
}

// FDataProviderConverter - тип функций сборки и конвертации данных для дата провайдеров
type FDataProviderConverter func() map[uint32]string

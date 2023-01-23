package tmpl

type Indexer struct {
	iType uint8 // ячейка в таблице. Полный адрес ячейка:строка:колонка
	idx   map[uint32]string
}

func NewCellIndexer() Indexer {
	return Indexer{
		iType: TYPE_INDEXER_TABLE_CELL,
		idx:   make(map[uint32]string),
	}
}

func NewRowIndexer() Indexer {
	return Indexer{
		iType: TYPE_INDEXER_LIST_ROW,
		idx:   make(map[uint32]string),
	}
}

func NewTextParamIndexer() Indexer {
	return Indexer{
		iType: TYPE_INDEXER_PARAM_IN_TEXT,
		idx:   make(map[uint32]string),
	}
}

func (i Indexer) AddData(data string, address ...uint32) {
	if !ApproveIndexerAddress(i.iType, address...) {
		return
	}

	completedAddress := MakeIndexerAddress(i.iType, address...)
	i.idx[completedAddress] = data
}

func (i Indexer) GetDataPartialAddress(address ...uint32) string {
	if !ApproveIndexerAddress(i.iType, address...) {
		return ""
	}

	completedAddress := MakeIndexerAddress(i.iType, address...)
	data, _ := i.idx[completedAddress]

	return data
}

func (i Indexer) GetDataFullAddress(address uint32) string {
	data, _ := i.idx[address]

	return data
}

package termui

type UpdatingDataProvider interface {
	UpdateData(data string, address ...uint32)
}

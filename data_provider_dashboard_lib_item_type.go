package termui

type DashboardLibItem struct {
	P, R, C  uint32
	iconView string
}

func NewDashboardLibsItem(iconView string, r, c, p uint32, styles ...Style) *DashboardLibItem {
	var str string

	if len(styles) == 0 {
		str = iconView
	} else {
		str = FormatStrWithStyle(iconView, styles[0])
	}

	return &DashboardLibItem{
		R:        r,
		C:        c,
		P:        p,
		iconView: str,
	}
}

func (dli *DashboardLibItem) String() string {
	return dli.iconView
}

func (dli *DashboardLibItem) Address() uint32 {
	return MakeDataProviderAddress(dli.P, dli.R, dli.C)
}

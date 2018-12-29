package orderbook

type level struct {
	LevelPrice  float64
	TotalVolume int64
	OrderCount  int64
	Orders      map[orderID]order
	parentPage  *page
}

func (lvl *level) exist(ord *order) bool {
	_, exist := lvl.Orders[ord.LocalID]

	return exist
}

func (lvl *level) mergeLevel(l *level) {
	lvl.TotalVolume += l.TotalVolume
}

func (lvl *level) build(volume int64) {
	lvl.TotalVolume = volume
}

func (lvl *level) modify(volume int64) {
	lvl.TotalVolume = volume
}

func (lvl *level) remove() {
	delete(lvl.parentPage.levelCache, lvl.LevelPrice)

	lvl.LevelPrice = 0
	lvl.TotalVolume = 0
	lvl.OrderCount = 0
}

func createLevel(price float64, parent *page) *level {
	lvl := level{LevelPrice: price, parentPage: parent}

	return &lvl
}

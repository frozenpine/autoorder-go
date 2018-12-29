package orderbook

import (
	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type volumeHeap struct {
	parent *level
	heap   []autoorder.OrderID
}

func (oh *volumeHeap) Len() int {
	return len(oh.heap)
}

func (oh *volumeHeap) Swap(i, j int) {
	oh.heap[i], oh.heap[j] = oh.heap[j], oh.heap[i]
}

func (oh *volumeHeap) Less(i, j int) bool {
	left, exist := oh.parent.Orders[oh.heap[i]]
	if !exist {
		panic("heap data mismatch with Orders cache.")
	}

	right, exist := oh.parent.Orders[oh.heap[j]]
	if !exist {
		panic("heap data mismatch with Orders cache.")
	}

	return left.Volume >= right.Volume
}

func (oh *volumeHeap) Push(h interface{}) {
	oid := h.(autoorder.OrderID)

	oh.heap = append(oh.heap, oid)
}

func (oh *volumeHeap) Pop() interface{} {
	count := len(oh.heap)
	last := oh.heap[count-1]
	oh.heap = oh.heap[:count-1]

	return last
}

type level struct {
	LevelPrice  float64
	TotalVolume int64
	Orders      map[autoorder.OrderID]*order
	sysIDMapper map[int64]autoorder.OrderID
	heap        volumeHeap
	parentPage  *page
}

func (lvl *level) exist(ord *order) bool {
	_, exist := lvl.Orders[ord.LocalID]

	return exist
}

func (lvl *level) Count() int {
	heapCount := lvl.heap.Len()
	orderCount := len(lvl.Orders)

	if heapCount == orderCount {
		return orderCount
	}

	panic("heap data mismatch with Orders cache.")
}

func (lvl *level) splitVolumes() {
	// todo: 自动拆单
}

func (lvl *level) modify(volume int64) {
	lvl.TotalVolume = volume

	lvl.splitVolumes()
}

func (lvl *level) remove() {
	if lvl.parentPage != nil {
		delete(lvl.parentPage.Levels, lvl.LevelPrice)
	}

	// todo: level中的委托处理
}

func createLevel(price float64, vol int64, parent *page) *level {
	lvl := level{LevelPrice: price, TotalVolume: vol, parentPage: parent, Orders: make(map[autoorder.OrderID]*order)}
	lvl.heap.parent = &lvl

	lvl.splitVolumes()

	return &lvl
}

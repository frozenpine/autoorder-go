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

func (lvl *level) Exist(ord *order) bool {
	return lvl.ExistLocalID(ord.LocalID)
}

func (lvl *level) ExistLocalID(id autoorder.OrderID) bool {
	_, exist := lvl.Orders[id]

	return exist
}

func (lvl *level) ExistSysID(id int64) bool {
	localID, exist := lvl.sysIDMapper[id]

	if !exist {
		return false
	}

	if lvl.ExistLocalID(localID) {
		return true
	}

	panic("sysIDMapper data mismatch with Orders cache.")
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

func (lvl *level) Modify(volume int64) {
	lvl.TotalVolume = volume

	lvl.splitVolumes()
}

func (lvl *level) Remove() {
	if lvl.parentPage != nil {
		delete(lvl.parentPage.Levels, lvl.LevelPrice)
	}

	// todo: level中的委托处理, 反向FAK对冲当前Level的Volume
}

func createLevel(price float64, vol int64, parent *page) *level {
	lvl := level{LevelPrice: price, TotalVolume: vol, parentPage: parent, Orders: make(map[autoorder.OrderID]*order)}
	lvl.heap.parent = &lvl

	lvl.splitVolumes()

	return &lvl
}

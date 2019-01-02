package orderbook

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

const (
	tinyVolumeCount  int   = 5
	hugeVolumeFactor int64 = 100
	tinyVolumeFactor int64 = 10
)

// volumeHeap 根据order.Volume排序的大顶堆
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
	LevelPrice     float64
	Orders         map[autoorder.OrderID]*order
	sysIDMapper    map[int64]autoorder.OrderID
	heap           volumeHeap
	maxVolPerOrder int64
	parentPage     *page
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

func (lvl *level) TotalVolume() int64 {
	var total int64
	for _, ord := range lvl.Orders {
		total += ord.Volume
	}
	return total
}

func (lvl *level) Count() int {
	heapCount := lvl.heap.Len()
	orderCount := len(lvl.Orders)

	if heapCount == orderCount {
		return orderCount
	}

	panic("heap data mismatch with Orders cache.")
}

// GetOrder 根据OrderLocalID获得order对象
func (lvl *level) GetOrder(id autoorder.OrderID) (*order, error) {
	ord, exist := lvl.Orders[id]

	if exist {
		return ord, nil
	}

	return nil, fmt.Errorf("Order not exist with localID[%d]", id)
}

func (lvl *level) pushOrder(vol int64) {
	ord := createOrder(lvl.LevelPrice, vol, lvl)

	lvl.Orders[ord.LocalID] = ord

	heap.Push(&lvl.heap, ord.LocalID)
}

func (lvl *level) popOrder() *order {
	orderID := heap.Pop(&lvl.heap).(autoorder.OrderID)

	ord, err := lvl.GetOrder(orderID)

	if err != nil {
		panic("heap data mismatch with Orders cache.")
	}

	delete(lvl.Orders, ord.LocalID)
	if ord.SysID > 0 {
		delete(lvl.sysIDMapper, ord.SysID)
	}

	return ord
}

func (lvl *level) splitVolumes(vol int64) {
	if !validateVolume(lvl.maxVolPerOrder) {
		return
	}

	oriTotal := lvl.TotalVolume()

	if oriTotal >= vol {
		return
	}

	remainedVol := oriTotal - vol
	var ordVol int64

	for remainedVol > 0 {
		r := rand.New(rand.NewSource(time.Now().Unix()))

		if lvl.Count() < tinyVolumeCount {
			// 随机数范围为[0, n), 随机数值+1以避免vol出现0值
			if vol >= lvl.maxVolPerOrder {
				ordVol = r.Int63n(hugeVolumeFactor) + 1
			} else {
				ordVol = r.Int63n(tinyVolumeFactor) + 1
			}
		} else {
			ordVol = lvl.maxVolPerOrder
		}

		if ordVol >= remainedVol {
			ordVol = remainedVol
		}

		lvl.pushOrder(ordVol)

		remainedVol -= ordVol
	}
}

func (lvl *level) Modify(volume int64) {
	diffVolume := lvl.TotalVolume() - volume

	if diffVolume == 0 {
		return
	}

	volRemained := diffVolume

	if diffVolume > 0 {
		// 新的Volume量少, 需要取消原有Level中委托的量
		for {
			maxVolOrder := lvl.popOrder()

			maxVolOrder.cancel()

			if maxVolOrder.Volume >= volRemained {
				// todo: 生成新的(maxVolOrder.Volume - volRemained)差额委托，push到level中
				break
			}

			volRemained -= maxVolOrder.Volume
		}
	} else {
		lvl.splitVolumes(volume)
	}
}

func (lvl *level) Remove() {
	if lvl.parentPage != nil {
		delete(lvl.parentPage.Levels, lvl.LevelPrice)
	}

	// todo: level中的委托处理, 反向FAK对冲当前Level的Volume
}

func createLevel(price float64, vol int64, parent *page) *level {
	lvl := level{
		LevelPrice: price,
		parentPage: parent,
		Orders:     make(map[autoorder.OrderID]*order)}
	lvl.heap.parent = &lvl

	if parent != nil && parent.parentBook != nil {
		lvl.maxVolPerOrder = parent.parentBook.MaxVolPerOrder
	}

	lvl.splitVolumes(vol)

	return &lvl
}

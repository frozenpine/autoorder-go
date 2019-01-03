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

func (h *volumeHeap) Len() int {
	return len(h.heap)
}

func (h *volumeHeap) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *volumeHeap) Less(i, j int) bool {
	left, exist := h.parent.Orders[h.heap[i]]
	if !exist {
		panic("heap data mismatch with Orders cache.")
	}

	right, exist := h.parent.Orders[h.heap[j]]
	if !exist {
		panic("heap data mismatch with Orders cache.")
	}

	return left.Volume >= right.Volume
}

func (h *volumeHeap) Push(v interface{}) {
	oid := v.(autoorder.OrderID)
	h.heap = append(h.heap, oid)
}

func (h *volumeHeap) Pop() interface{} {
	count := len(h.heap)
	last := h.heap[count-1]
	h.heap = h.heap[:count-1]

	return last
}

func (h *volumeHeap) removeAt(idx int) {
	if idx <= h.Len()-2 {
		h.heap = append(h.heap[:idx], h.heap[idx+1:]...)
	} else {
		h.heap = h.heap[:idx]
	}

	heap.Init(h)
}

type orderAPI interface {
	Order(price float64, vol int64) (autoorder.OrderID, error)
	Cancel(oid autoorder.OrderID) error
	Hedge(price float64, vol int64) error
}

type level struct {
	LevelPrice     float64
	Orders         map[autoorder.OrderID]*order
	sysIDMapper    map[int64]autoorder.OrderID
	heap           volumeHeap
	maxVolPerOrder int64
	api            orderAPI
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

func (lvl *level) NewOrder(vol int64) error {
	oid, err := lvl.api.Order(lvl.LevelPrice, vol)

	if err != nil {
		return err
	}

	ord := newOrder(vol, oid)

	lvl.Orders[oid] = ord

	heap.Push(&lvl.heap, oid)

	return nil
}

// GetOrder 根据OrderLocalID获得order对象
func (lvl *level) GetOrder(oid autoorder.OrderID) (*order, error) {
	ord, exist := lvl.Orders[oid]

	if exist {
		return ord, nil
	}

	return nil, fmt.Errorf("Order not exist with localID[%d]", oid)
}

func (lvl *level) DeleteOrder(oid autoorder.OrderID) (*order, error) {
	ord, err := lvl.GetOrder(oid)

	if err != nil {
		return nil, err
	}

	delete(lvl.Orders, ord.LocalID)
	delete(lvl.sysIDMapper, ord.SysID)

	for i := 0; i < lvl.heap.Len(); i++ {
		if lvl.heap.heap[i] != ord.LocalID {
			continue
		}

		lvl.heap.removeAt(i)
		break
	}

	return ord, nil
}

func (lvl *level) PeekOrder() *order {
	oid := lvl.heap.heap[0]

	ord, err := lvl.GetOrder(oid)

	if err != nil {
		panic("heap data mismatch with Orders cache.")
	}

	return ord
}

func (lvl *level) PopOrder() *order {
	oid := heap.Pop(&lvl.heap).(autoorder.OrderID)

	ord, err := lvl.GetOrder(oid)

	if err != nil {
		panic("heap data mismatch with Orders cache.")
	}

	return ord
}

func (lvl *level) splitVolumes(vol int64) {
	oriTotal := lvl.TotalVolume()

	if oriTotal >= vol {
		return
	}

	if !validateVolume(lvl.maxVolPerOrder) {
		lvl.NewOrder(vol)
		return
	}

	remainedVol := vol - oriTotal
	var ordVol int64

	for remainedVol > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

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

		if ordVol > remainedVol {
			ordVol = remainedVol
		}

		lvl.NewOrder(ordVol)

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
			maxVolOrder := lvl.PeekOrder()

			// maxVolOrder.cancel()

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
	lvl.api.Hedge(lvl.LevelPrice, lvl.TotalVolume())
}

func newLevel(price float64, vol int64, maxVol int64, api orderAPI) *level {
	lvl := level{
		LevelPrice:     price,
		Orders:         make(map[autoorder.OrderID]*order),
		maxVolPerOrder: maxVol,
		api:            api}
	lvl.heap.parent = &lvl

	lvl.splitVolumes(vol)

	return &lvl
}

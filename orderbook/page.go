package orderbook

import (
	"container/heap"
	"errors"
)

type priceHeap struct {
	parent *page
	heap   []float64
}

func (hp *priceHeap) isMinHeap() bool {
	return hp.parent.Direction == Sell
}

func (hp *priceHeap) Len() int {
	return len(hp.heap)
}

func (hp *priceHeap) Swap(i, j int) {
	hp.heap[i], hp.heap[j] = hp.heap[j], hp.heap[i]
}

func (hp *priceHeap) Less(i, j int) bool {
	if hp.isMinHeap() {
		return hp.heap[i] < hp.heap[j]
	}
	return hp.heap[i] > hp.heap[j]
}

func (hp *priceHeap) Push(h interface{}) {
	lvlPrice := h.(float64)

	if lvlPrice == 0 {
		return
	}

	hp.heap = append(hp.heap, lvlPrice)
}

func (hp *priceHeap) Pop() (x interface{}) {
	count := len(hp.heap)
	last := hp.heap[count-1]
	hp.heap = hp.heap[:count-1]

	return last
}

func (hp *priceHeap) peek(idx int) float64 {
	return hp.heap[idx]
}

func (hp *priceHeap) removeAt(idx int) {
	if idx <= hp.Len()-2 {
		hp.heap = append(hp.heap[:idx], hp.heap[idx+1:]...)
	} else {
		hp.heap = hp.heap[:idx]
	}

	heap.Init(hp)
}

type page struct {
	Direction  direction
	parentBook *Book
	levelCache map[float64]*level
	levelHeap  priceHeap
}

// BestPrice 获取当前方向上的挂单最优价
func (p *page) BestPrice() float64 {
	price := p.levelHeap.heap[0]

	if _, exist := p.levelCache[price]; exist {
		return price
	}

	panic("priceHeap mismatch with levelCache.")
}

// BestLevel 获取当前方向上的最优挂单Level
func (p *page) BestLevel() *level {
	lvl, _ := p.levelCache[p.BestPrice()]

	return lvl
}

// GetLevel 获取当前方向上特定价格Level
func (p *page) GetLevel(price float64) (*level, error) {
	if lvl, exist := p.levelCache[price]; exist {
		return lvl, nil
	}

	return nil, errors.New("level not exist")
}

// Size 获取当前方向上的挂单Level数
func (p *page) Size() int {
	heapLen := p.levelHeap.Len()
	cacheLen := len(p.levelCache)

	if heapLen == cacheLen {
		return heapLen
	}

	panic("priceHeap size mismatch with levelCache.")
}

// PopLevel 删除当前方向上的最优价Level
func (p *page) PopLevel() *level {
	lvlPrice := heap.Pop(&p.levelHeap).(float64)

	lvl, exist := p.levelCache[lvlPrice]
	defer lvl.remove()

	if !exist {
		panic("priceHeap data mismatch with levelCache.")
	}

	return lvl
}

// AddLevel 在当前方向上新增一个价格Level
func (p *page) AddLevel(price float64, volume int64) bool {
	if _, err := p.GetLevel(price); err == nil {
		return false
	}

	defer heap.Push(&p.levelHeap, price)

	newLevel := createLevel(price, p)
	newLevel.build(volume)
	p.levelCache[price] = newLevel

	return true
}

// RemoveLevel 在当前方向上删除价格Level
func (p *page) RemoveLevel(lvlPrice float64) *level {
	lvl, exist := p.levelCache[lvlPrice]

	if !exist {
		return nil
	}

	defer lvl.remove()

	for i := 0; i < p.levelHeap.Len(); i++ {
		if p.levelHeap.peek(i) != lvlPrice {
			continue
		}

		p.levelHeap.removeAt(i)
		break
	}

	return lvl
}

// ModifyLevel 在当前方向上修改对应价格Level的量
func (p *page) ModifyLevel(price float64, volume int64) bool {
	lvl, err := p.GetLevel(price)

	if err != nil {
		return false
	}

	lvl.modify(volume)

	return true
}

func createPage(d direction, parent *Book) *page {
	p := page{Direction: d, parentBook: parent, levelCache: make(map[float64]*level)}
	hp := priceHeap{parent: &p, heap: make([]float64, 0, 10)}
	p.levelHeap = hp

	return &p
}

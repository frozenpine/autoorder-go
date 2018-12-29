package orderbook

import (
	"container/heap"
	"errors"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type priceHeap struct {
	asc  bool
	heap []float64
}

func (hp *priceHeap) Len() int {
	return len(hp.heap)
}

func (hp *priceHeap) Swap(i, j int) {
	hp.heap[i], hp.heap[j] = hp.heap[j], hp.heap[i]
}

func (hp *priceHeap) Less(i, j int) bool {
	if hp.asc {
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

func (hp *priceHeap) Pop() interface{} {
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
	direction  autoorder.Direction
	parentBook *Book
	Levels     map[float64]*level
	heap       priceHeap
}

// Overlapped 判断价格是否和当前方向上重叠
func (p *page) Overlapped(price float64) bool {
	switch p.direction {
	case autoorder.Buy:
		return price <= p.BestPrice()
	case autoorder.Sell:
		return price >= p.BestPrice()
	default:
		return true
	}
}

// BestPrice 获取当前方向上的挂单最优价
func (p *page) BestPrice() float64 {
	price := p.heap.heap[0]

	if _, exist := p.Levels[price]; exist {
		return price
	}

	panic("heap data mismatch with Levels cache.")
}

// BestLevel 获取当前方向上的最优挂单Level
func (p *page) BestLevel() *level {
	lvl, _ := p.Levels[p.BestPrice()]

	return lvl
}

// GetLevel 获取当前方向上特定价格Level
func (p *page) GetLevel(price float64) (*level, error) {
	if lvl, exist := p.Levels[price]; exist {
		return lvl, nil
	}

	return nil, errors.New("level not exist")
}

// Size 获取当前方向上的挂单Level数
func (p *page) Size() int {
	heapLen := p.heap.Len()
	cacheLen := len(p.Levels)

	if heapLen == cacheLen {
		return cacheLen
	}

	panic("heap size mismatch with Levels cache.")
}

// PopLevel 删除当前方向上的最优价Level
func (p *page) PopLevel() *level {
	lvlPrice := heap.Pop(&p.heap).(float64)

	lvl, exist := p.Levels[lvlPrice]
	defer lvl.remove()

	if !exist {
		panic("heap data mismatch with Levels Cache.")
	}

	return lvl
}

// AddLevel 在当前方向上新增一个价格Level
func (p *page) AddLevel(price float64, volume int64) bool {
	if _, err := p.GetLevel(price); err == nil {
		return false
	}

	defer heap.Push(&p.heap, price)

	newLevel := createLevel(price, volume, p)
	p.Levels[price] = newLevel

	return true
}

// RemoveLevel 在当前方向上删除价格Level
func (p *page) RemoveLevel(lvlPrice float64) *level {
	lvl, exist := p.Levels[lvlPrice]

	if !exist {
		return nil
	}

	defer lvl.remove()

	for i := 0; i < p.heap.Len(); i++ {
		if p.heap.peek(i) != lvlPrice {
			continue
		}

		p.heap.removeAt(i)
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

func createPage(d autoorder.Direction, parent *Book) *page {
	p := page{direction: d, parentBook: parent, Levels: make(map[float64]*level)}
	switch d {
	case autoorder.Buy:
		p.heap = priceHeap{asc: false, heap: make([]float64, 0, 10)}
	case autoorder.Sell:
		p.heap = priceHeap{asc: true, heap: make([]float64, 0, 10)}
	default:
		panic("Invalid direction")
	}

	return &p
}

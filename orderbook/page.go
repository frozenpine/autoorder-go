package orderbook

import (
	"container/heap"
	"errors"
)

type page struct {
	Direction  direction
	parentBook *Book
	levelCache map[float64]*level
	priceHeap  []float64
}

func (p *page) isMinHeap() bool {
	return p.Direction == Sell
}

func (p *page) bestPrice() float64 {
	price := p.priceHeap[0]
	if _, exist := p.levelCache[price]; exist {
		return price
	}

	panic("priceHeap mismatch with levelCache.")
}

func (p *page) bestLevel() *level {
	lvl, _ := p.levelCache[p.bestPrice()]

	return lvl
}

func (p *page) getLevel(price float64) (*level, error) {
	if lvl, exist := p.levelCache[price]; exist {
		return lvl, nil
	}

	return nil, errors.New("level not exist")
}

func (p *page) size() int {
	heapLen := len(p.priceHeap)
	cacheLen := len(p.levelCache)

	if heapLen == cacheLen {
		return heapLen
	}

	panic("priceHeap size mismatch with levelCache.")
}

func (p *page) popLevel() *level {
	lvlPrice := heap.Pop(p).(float64)

	lvl, exist := p.levelCache[lvlPrice]
	delete(p.levelCache, lvlPrice)

	if !exist {
		panic("priceHeap data mismatch with levelCache.")
	}

	return lvl
}

func (p *page) addLevel(price float64, volume int64) bool {
	if _, err := p.getLevel(price); err == nil {
		return false
	}

	heap.Push(p, price)

	newLevel := createLevel(price, p)
	newLevel.build(volume)
	p.levelCache[price] = newLevel

	return true
}

func (p *page) removeLevel(lvlPrice float64) *level {
	lvl, exist := p.levelCache[lvlPrice]
	if !exist {
		return nil
	}

	defer heap.Init(p)

	delete(p.levelCache, lvlPrice)

	oriLen := len(p.priceHeap)

	for i := 0; i < oriLen; i++ {
		if p.priceHeap[i] != lvlPrice {
			continue
		}

		if i <= oriLen-2 {
			p.priceHeap = append(p.priceHeap[:i], p.priceHeap[i+1:]...)
		} else {
			p.priceHeap = p.priceHeap[:i]
		}
		break
	}

	return lvl
}

func (p *page) modifyLevel(price float64, volume int64) bool {
	lvl, err := p.getLevel(price)

	if err != nil {
		return false
	}

	lvl.modify(volume)

	return true
}

func (p *page) Len() int {
	return len(p.priceHeap)
}

func (p *page) Swap(i, j int) {
	p.priceHeap[i], p.priceHeap[j] = p.priceHeap[j], p.priceHeap[i]
}

func (p *page) Less(i, j int) bool {
	if p.isMinHeap() {
		return p.priceHeap[i] < p.priceHeap[j]
	}
	return p.priceHeap[i] > p.priceHeap[j]
}

func (p *page) Push(h interface{}) {
	lvlPrice := h.(float64)

	if lvlPrice == 0 {
		return
	}

	if _, exist := p.levelCache[lvlPrice]; exist {
		panic("conflict level price")
	}

	p.priceHeap = append(p.priceHeap, lvlPrice)
}

func (p *page) Pop() (x interface{}) {
	count := len(p.priceHeap)
	last := p.priceHeap[count-1]
	p.priceHeap = p.priceHeap[:count-1]

	return last
}

func createPage(d direction, parent *Book) *page {
	p := page{Direction: d, parentBook: parent, priceHeap: make([]float64, 0, 10), levelCache: make(map[float64]*level)}

	return &p
}

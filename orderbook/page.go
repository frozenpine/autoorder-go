package orderbook

import (
	"container/heap"
	"errors"
)

// type pageIterator struct {
// 	pageData *page
// 	index    int
// }

// func (i *pageIterator) HasNext() bool {
// 	return i.index < i.pageData.Len()
// }

// func (i *pageIterator) nextIndex() {
// 	if i.pageData.isMinHeap() {
// 		i.index++
// 	} else {
// 		i.index++
// 	}
// }

// func (i *pageIterator) Next() *level {
// 	levelPrice := i.pageData.priceHeap[i.index]

// 	defer i.nextIndex()

// 	if lvl, err := i.pageData.getLevel(levelPrice); err != nil {
// 		panic(fmt.Sprintf("Error occoured during iteration: %s", err.Error()))
// 	} else {
// 		return lvl
// 	}
// }

type page struct {
	Direction  direction
	parentBook *Book
	levelCache map[float64]level
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

	return &lvl
}

func (p *page) getLevel(price float64) (*level, error) {
	if lvl, exist := p.levelCache[price]; exist {
		return &lvl, nil
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

func (p *page) addLevel(price float64, volume int64) bool {
	if _, err := p.getLevel(price); err == nil {
		return false
	}

	newLevel := level{LevelPrice: price}
	defer newLevel.build(volume)

	heap.Push(p, newLevel)

	return true
}

func (p *page) modifyLevel(price float64, volume int64) bool {
	lvl, err := p.getLevel(price)

	if err != nil {
		return false
	}

	lvl.modify(volume)

	return true
}

// func (p *page) Iterator() *pageIterator {
// 	return &pageIterator{pageData: p, index: 0}
// }

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
	newLevel := h.(level)

	if newLevel.LevelPrice == 0 {
		return
	}

	if p.levelCache == nil {
		p.levelCache = make(map[float64]level)
	}

	if p.priceHeap == nil {
		p.priceHeap = make([]float64, 0, 10)
	}

	if oriLevel, exist := p.levelCache[newLevel.LevelPrice]; exist {
		oriLevel.mergeLevel(newLevel)
	} else {
		p.levelCache[newLevel.LevelPrice] = newLevel
		p.priceHeap = append(p.priceHeap, newLevel.LevelPrice)
	}
}

func (p *page) Pop() (x interface{}) {
	count := len(p.priceHeap)
	last := p.priceHeap[count-1]
	p.priceHeap = p.priceHeap[:count-1]

	lvl, _ := p.levelCache[last]
	delete(p.levelCache, last)

	return lvl
}

func createPage(d direction, parent *Book) *page {
	p := page{Direction: d, parentBook: parent}

	return &p
}

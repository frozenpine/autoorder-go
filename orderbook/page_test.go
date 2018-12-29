package orderbook

import (
	"container/heap"
	"testing"
)

func TestPageHeap(t *testing.T) {
	buyPage := new(page)
	sellPage := new(page)

	buyPage.Direction = Buy
	sellPage.Direction = Sell

	heap.Init(buyPage)
	heap.Init(sellPage)

	for price := 2; price < 10; price++ {
		heap.Push(buyPage, level{LevelPrice: float64(price)})
		heap.Push(sellPage, level{LevelPrice: float64(price)})
	}

	if x := buyPage.bestPrice(); x != 9 {
		t.Error("Buy page error")
	}

	if x := sellPage.bestPrice(); x != 2 {
		t.Error("Sell page error")
	}

	heap.Push(buyPage, level{LevelPrice: float64(6)})
	heap.Push(sellPage, level{LevelPrice: float64(1)})

	if buyPage.size() != 8 {
		t.Error("Buy page push error.")
	} else if buyPage.bestLevel().LevelPrice != 9 {
		t.Error("Buy page sort error.")
	} else {
		t.Log(buyPage.priceHeap)
	}

	if sellPage.size() != 9 {
		t.Error("Sell page push error.")
	} else if sellPage.bestPrice() != 1 {
		t.Error("Sell page sort error.")
	} else {
		t.Log(sellPage.priceHeap)
	}

	sell := heap.Pop(sellPage).(level)
	if sell.LevelPrice != 1 {
		t.Error("Sell page pop error.")
	} else {
		t.Log("Poped sell price:", sell.LevelPrice)
		// for iter := sellPage.Iterator(); iter.HasNext(); {
		// 	t.Log(iter.Next().LevelPrice)
		// }
	}

	buy := heap.Pop(buyPage).(level)
	if buy.LevelPrice != 9 {
		t.Error("Buy page pop error.")
	} else {
		t.Log("Poped buy price:", buy.LevelPrice)
		// for iter := buyPage.Iterator(); iter.HasNext(); {
		// 	t.Log(iter.Next().LevelPrice)
		// }
	}
}

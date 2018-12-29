package orderbook

import (
	"testing"
)

func TestPageHeap(t *testing.T) {
	buyPage := createPage(Buy, nil)
	sellPage := createPage(Sell, nil)

	for price := 2; price < 10; price++ {
		buyPage.AddLevel(float64(price), 0)
		sellPage.AddLevel(float64(price), 0)
	}

	if x := buyPage.BestPrice(); x != 9 {
		t.Error("Buy page error")
	}

	if x := sellPage.BestPrice(); x != 2 {
		t.Error("Sell page error")
	}

	if ok := buyPage.AddLevel(6, 0); ok || buyPage.Size() != 8 {
		t.Error("Buy page push error.")
	} else if buyPage.BestLevel().LevelPrice != 9 {
		t.Error("Buy page sort error.")
	} else {
		t.Log(buyPage.levelHeap)
	}

	if ok := sellPage.AddLevel(1, 0); !ok || sellPage.Size() != 9 {
		t.Error("Sell page push error.")
	} else if sellPage.BestPrice() != 1 {
		t.Error("Sell page sort error.")
	} else {
		t.Log(sellPage.levelHeap)
	}

	sell := sellPage.PopLevel()
	if sell.LevelPrice != 1 {
		t.Error("Sell page pop error.")
	} else if lvl := sellPage.RemoveLevel(6); lvl == nil || sellPage.Size() != 7 || sellPage.BestPrice() != 2 {
		t.Error("Sell page remove level error.")
	} else {
		t.Log(sellPage.levelCache)
		t.Log(sellPage.levelHeap)
	}

	buy := buyPage.PopLevel()
	if buy.LevelPrice != 9 {
		t.Error("Buy page pop error.")
	} else if lvl := buyPage.RemoveLevel(8); lvl == nil || buyPage.Size() != 6 || buyPage.BestPrice() != 7 {
		t.Error("Buy page remove level error.")
	} else {
		t.Log(buyPage.levelCache)
		t.Log(buyPage.levelHeap)
	}
}

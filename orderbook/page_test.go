package orderbook

import (
	"testing"
)

func TestPageHeap(t *testing.T) {
	buyPage := createPage(Buy, nil)
	sellPage := createPage(Sell, nil)

	for price := 2; price < 10; price++ {
		buyPage.addLevel(float64(price), 0)
		sellPage.addLevel(float64(price), 0)
	}

	if x := buyPage.bestPrice(); x != 9 {
		t.Error("Buy page error")
	}

	if x := sellPage.bestPrice(); x != 2 {
		t.Error("Sell page error")
	}

	if ok := buyPage.addLevel(6, 0); ok || buyPage.size() != 8 {
		t.Error("Buy page push error.")
	} else if buyPage.bestLevel().LevelPrice != 9 {
		t.Error("Buy page sort error.")
	} else {
		t.Log(buyPage.priceHeap)
	}

	if ok := sellPage.addLevel(1, 0); !ok || sellPage.size() != 9 {
		t.Error("Sell page push error.")
	} else if sellPage.bestPrice() != 1 {
		t.Error("Sell page sort error.")
	} else {
		t.Log(sellPage.priceHeap)
	}

	sell := sellPage.popLevel()
	if sell.LevelPrice != 1 {
		t.Error("Sell page pop error.")
	} else if lvl := sellPage.removeLevel(6); lvl == nil || sellPage.size() != 7 || sellPage.bestPrice() != 2 {
		t.Error("Sell page remove level error.")
	} else {
		t.Log(sellPage.levelCache)
		t.Log(sellPage.priceHeap)
	}

	buy := buyPage.popLevel()
	if buy.LevelPrice != 9 {
		t.Error("Buy page pop error.")
	} else if lvl := buyPage.removeLevel(8); lvl == nil || buyPage.size() != 6 || buyPage.bestPrice() != 7 {
		t.Error("Buy page remove level error.")
	} else {
		t.Log(buyPage.levelCache)
		t.Log(buyPage.priceHeap)
	}
}

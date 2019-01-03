package orderbook

import (
	"testing"

	"gitlab.quantdo.cn/yuanyang/autoorder/trader"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

func TestPageHeap(t *testing.T) {
	mock := new(trader.MockTrader)
	buyPage := newPage(autoorder.Buy, 1000, mock)
	sellPage := newPage(autoorder.Sell, 0, mock)

	for price := 2; price < 10; price++ {
		buyPage.AddLevel(float64(price), 10)
		sellPage.AddLevel(float64(price), 10)
	}

	if x := buyPage.BestPrice(); x != 9 {
		t.Error("Buy page error")
	}

	if x := sellPage.BestPrice(); x != 2 {
		t.Error("Sell page error")
	}

	if ok := buyPage.AddLevel(6, 10); ok || buyPage.Size() != 8 {
		t.Error("Buy page push error.")
	} else if buyPage.BestLevel().LevelPrice != 9 {
		t.Error("Buy page sort error.")
	} else {
		t.Log(buyPage.heap)
	}

	if ok := sellPage.AddLevel(1, 10); !ok || sellPage.Size() != 9 {
		t.Error("Sell page push error.")
	} else if sellPage.BestPrice() != 1 {
		t.Error("Sell page sort error.")
	} else {
		t.Log(sellPage.heap)
	}

	sell := sellPage.PopLevel()
	if sell.LevelPrice != 1 {
		t.Error("Sell page pop error.")
	} else if lvl := sellPage.RemoveLevel(6); lvl == nil || sellPage.Size() != 7 || sellPage.BestPrice() != 2 {
		t.Error("Sell page remove level error.")
	} else {
		t.Log(sellPage.Levels)
		t.Log(sellPage.heap)
	}

	buy := buyPage.PopLevel()
	if buy.LevelPrice != 9 {
		t.Error("Buy page pop error.")
	} else if lvl := buyPage.RemoveLevel(8); lvl == nil || buyPage.Size() != 6 || buyPage.BestPrice() != 7 {
		t.Error("Buy page remove level error.")
	} else {
		t.Log(buyPage.Levels)
		t.Log(buyPage.heap)
	}
}

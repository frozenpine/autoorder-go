package orderbook

import (
	"testing"

	"gitlab.quantdo.cn/yuanyang/autoorder"
	"gitlab.quantdo.cn/yuanyang/autoorder/trader"
)

func TestCalculateBlock(t *testing.T) {
	mock := new(trader.MockTrader)
	ob := CreateOrderBook("SHFE", "fu1905", 1000, 0.1, 0.1, mock)

	ob.UpdateBlock(autoorder.Buy, 123)
	ob.UpdateBlock(autoorder.Sell, 123)

	if ob.ceilBlock != nil || ob.floorBlock != nil {
		t.Error("UpdateBlock fail with MakeBlock false")
	}

	ob.MakeBlock = true
	ob.UpdateBlock(autoorder.Sell, 321)
	ob.UpdateBlock(autoorder.Buy, 123)

	if ob.ceilBlock == nil || ob.floorBlock == nil {
		t.Error("UpdateBlock fail with MakeBlock true")
	}
	if ob.ceilBlock.LevelPrice < 321+0.1*10 {
		t.Error("Calculate ceil price fail.")
	} else {
		t.Logf("Ceil block: %d@%f", ob.ceilBlock.TotalVolume(), ob.ceilBlock.LevelPrice)
	}
	if ob.floorBlock.LevelPrice > 123-0.1*10 {
		t.Error("Calculate floor price fail.")
	} else {
		t.Logf("Floor block: %d@%f", ob.floorBlock.TotalVolume(), ob.floorBlock.LevelPrice)
	}

	ob.StopPrice = 121
	ob.LimitPrice = 323
	ob.TickPrice = 1
	ob.UpdateBlock(autoorder.Buy, 123)
	ob.UpdateBlock(autoorder.Sell, 321)

	if ob.ceilBlock.LevelPrice > 323 || ob.floorBlock.LevelPrice < 121 {
		t.Error("Calculate block price fail.")
	} else {
		t.Logf("Ceil block: %d@%f", ob.ceilBlock.TotalVolume(), ob.ceilBlock.LevelPrice)
		t.Logf("Floor block: %d@%f", ob.floorBlock.TotalVolume(), ob.floorBlock.LevelPrice)
	}

	ob.StopPrice = 113
	ob.LimitPrice = 331
	ob.UpdateBlock(autoorder.Buy, 123)
	ob.UpdateBlock(autoorder.Sell, 321)

	if ob.floorBlock.LevelPrice != 115 || ob.ceilBlock.LevelPrice != 329 {
		t.Error("Calculate block price fail.")
	} else {
		t.Logf("Ceil block: %d@%f", ob.ceilBlock.TotalVolume(), ob.ceilBlock.LevelPrice)
		t.Logf("Floor block: %d@%f", ob.floorBlock.TotalVolume(), ob.floorBlock.LevelPrice)
	}
}

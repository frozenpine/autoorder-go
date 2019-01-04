package orderbook

import (
	"encoding/json"
	"testing"

	"gitlab.quantdo.cn/yuanyang/autoorder/trader"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

func TestCreateOb(t *testing.T) {
	mock := new(trader.MockTrader)
	ob := CreateOrderBook("SHFE", "fu1905", 1000, 0.1, 0.1, mock)

	if id := ob.Identity(); id != "SHFE.fu1905" {
		t.Error("Identity fail")
	} else {
		t.Log(id)
	}

	if ob.Asks.direction != autoorder.Sell || ob.Bids.direction != autoorder.Buy {
		t.Error("Page directiion fail.")
	}
}

func TestJsonMarshmal(t *testing.T) {
	mock := new(trader.MockTrader)
	ob := CreateOrderBook("SHFE", "fu1905", 1000, 0.1, 0.1, mock)

	data, _ := json.Marshal(ob.Snapshot())
	t.Log(string(data))
}

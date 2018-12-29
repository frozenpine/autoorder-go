package orderbook

import (
	"testing"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

func TestCreateOb(t *testing.T) {
	ob := CreateOrderBook("SHFE", "fu1905", 1000, 0.1, 0.1, nil)

	if id := ob.Identity(); id != "SHFE.fu1905" {
		t.Error("Identity fail")
	} else {
		t.Log(id)
	}

	if ob.Asks.direction != autoorder.Sell || ob.Bids.direction != autoorder.Buy {
		t.Error("Page directiion fail.")
	}
}

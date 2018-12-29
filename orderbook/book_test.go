package orderbook

import (
	"testing"
)

func TestCreateOb(t *testing.T) {
	ob := CreateOrderBook("SHFE", "fu1905", 1000, 0.1, nil)

	if id := ob.Identity(); id != "SHFE.fu1905" {
		t.Error("Identity fail")
	} else {
		t.Log(id)
	}

	if ob.Asks.Direction != Sell || ob.Bids.Direction != Buy {
		t.Error("Page directiion fail.")
	}

	t.Log(ob.Asks.Size())
}

package orderbook

import (
	"testing"

	"gitlab.quantdo.cn/yuanyang/autoorder"
	"gitlab.quantdo.cn/yuanyang/autoorder/trader"
)

func TestLevelHeap(t *testing.T) {
	mock := new(trader.MockTrader)
	page := newPage(autoorder.Buy, 1000, mock)

	level := newLevel(12.6, 100, page, true)

	if level.TotalVolume() != 100 {
		t.Error("Split volume error.")
	}

	oriCount := level.Count()

	peekOrder := level.PeekOrder()

	t.Log(peekOrder.Volume)
	for _, oid := range level.heap.heap[1:] {
		ord, err := level.GetOrder(oid)
		if err != nil {
			t.Error(err)
		}
		if peekOrder.Volume < ord.Volume {
			t.Error("volumeHeep sort fail.")
		} else {
			t.Log(ord.Volume)
		}
	}

	if level.Count() != oriCount {
		t.Error("level PeekOrder failed.")
	} else if !level.Exist(peekOrder) {
		t.Error("Exist func fail.")
	} else if !level.ExistLocalID(peekOrder.LocalID) {
		t.Error("ExistLocalID func fail.")
	} else {
		t.Log(level.Snapshot())
	}

	popOrder := level.PopOrder()

	if popOrder.LocalID != peekOrder.LocalID {
		t.Error("PeekOrder mismatch with PopOrder")
	} else if level.Count() != oriCount-1 || level.ExistLocalID(popOrder.LocalID) {
		t.Error("PopOrder failed.")
	} else {
		t.Log(level.Snapshot())
	}

	popedVolume := level.TotalVolume()

	if popedVolume != 100-popOrder.Volume {
		t.Error("TotalVolume func fail")
	}

	if level.Modify(popedVolume - 10) {
		if level.TotalVolume() != popedVolume-10 {
			t.Error("Modify down func fail.")
		} else {
			t.Log(level.Count())
			t.Log(level.Snapshot())
		}
	}

	level.Modify(1001)

	if level.Count() < 6 {
		t.Error("Modify up func fail.")
	} else {
		t.Log(level.Snapshot())
	}
}

package orderbook

import (
	"sync/atomic"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type priceItem struct {
	Price  float64
	Volume int64
}

type order struct {
	priceItem
	identity
	LocalID     autoorder.OrderID
	SysID       int64
	parentLevel *level
}

func (ord *order) cancel() {
	if !ord.parentLevel.exist(ord) {
		return
	}

	delete(ord.parentLevel.Orders, ord.LocalID)

	atomic.AddInt64(&ord.parentLevel.TotalVolume, -ord.Volume)

	// todo: cancel order to trading system
}

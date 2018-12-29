package orderbook

import (
	"sync/atomic"
)

type orderID int64

type order struct {
	priceItem
	identity
	LocalID     orderID
	SysID       int64
	parentLevel *level
}

func (ord *order) cancel() {
	if !ord.parentLevel.exist(ord) {
		return
	}

	delete(ord.parentLevel.Orders, ord.LocalID)

	atomic.AddInt64(&ord.parentLevel.OrderCount, -1)
	atomic.AddInt64(&ord.parentLevel.TotalVolume, -ord.Volume)

	// todo: cancel order to trading system
}

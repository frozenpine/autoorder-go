package orderbook

import (
	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type order struct {
	Volume       int64
	TradedVolume int64
	LocalID      autoorder.OrderID
	SysID        int64
}

func newOrder(vol int64, oid autoorder.OrderID) *order {
	ord := order{Volume: vol, LocalID: oid}

	return &ord
}

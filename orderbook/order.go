package orderbook

import (
	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type order struct {
	Volume       int64             `json:"Volume"`
	TradedVolume int64             `json:"TradedVolume"`
	LocalID      autoorder.OrderID `json:"OrderLocalID"`
	SysID        int64             `json:"OrderSysID"`
}

func newOrder(vol int64, oid autoorder.OrderID) *order {
	ord := order{Volume: vol, LocalID: oid}

	return &ord
}

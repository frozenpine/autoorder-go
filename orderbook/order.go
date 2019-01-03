package orderbook

import (
	"gitlab.quantdo.cn/yuanyang/autoorder"
)

// type priceItem struct {
// 	Price  float64
// 	Volume int64
// }

type orderStatus byte

const (
	// OnTheWay 委托在途
	OnTheWay orderStatus = 0x10
	// Accepted 委托已接受
	Accepted orderStatus = 0x20
	// Rejected 委托已拒绝
	Rejected orderStatus = 0x40
	// Queue 队列中
	Queue orderStatus = 1 << iota
	// PartTrade 部分成交
	PartTrade
	// AllTrade 全部成交
	AllTrade
	// Cancel 撤销
	Cancel
)

func (status orderStatus) ChangeStatus(s orderStatus) orderStatus {
	return s
}

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

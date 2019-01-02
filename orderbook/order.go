package orderbook

import (
	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type priceItem struct {
	Price  float64
	Volume int64
}

type order struct {
	priceItem
	LocalID     autoorder.OrderID
	SysID       int64
	parentLevel *level
}

func (ord *order) cancel() {
	if !ord.parentLevel.Exist(ord) {
		return
	}

	delete(ord.parentLevel.Orders, ord.LocalID)

	// todo: cancel order to trading system
}

func createOrder(price float64, vol int64, parent *level) *order {
	ord := order{
		priceItem:   priceItem{Price: price, Volume: vol},
		parentLevel: parent}

	return &ord
}

package autoorder

// TraderAPI autoorder通用报单接口
type TraderAPI interface {
	Order(d Direction, price float64, vol int64) (OrderID, error)
	FAK(d Direction, price float64, vol int64) (OrderID, error)
	Cancel(localID OrderID) error
}

// MarketAPI autoorder通用行情接口
type MarketAPI interface {
	Subscribe(ins ...string)
	UnSubscribe(ins ...string)
}

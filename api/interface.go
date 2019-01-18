package api

import (
	"time"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type connAPI interface {
	Connect(addr string) error
	Disconnect() error
}

type loginAPI interface {
	Login(loginInfo map[string]string) error
	Logout() error
}

// TraderAPI autoorder通用报单接口
type TraderAPI interface {
	connAPI
	loginAPI
	Order(d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error)
	FAK(d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error)
	Cancel(localID autoorder.OrderID) error
	QueryOrders(instrumentID string, from, to time.Time) error
	QueryTrades() error
}

// MarketAPI autoorder通用行情接口
type MarketAPI interface {
	connAPI
	loginAPI
	Subscribe(ins ...string)
	UnSubscribe(ins ...string)
	SubscribeTopic(topic ...int)
	UnSubscribeTopic(topic ...int)
}

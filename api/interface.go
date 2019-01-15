package api

import (
	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type loginAPI interface {
	Login(loginInfo map[string]string) error
	Logout() error
}

// TraderAPI autoorder通用报单接口
type TraderAPI interface {
	loginAPI
	Order(d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error)
	FAK(d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error)
	Cancel(localID autoorder.OrderID) error
}

// MarketAPI autoorder通用行情接口
type MarketAPI interface {
	loginAPI
	Subscribe(ins ...string)
	UnSubscribe(ins ...string)
	SubscribeTopic(topic ...int)
	UnSubscribeTopic(topic ...int)
}

package trader

import (
	"fmt"
	"log"
	"sync/atomic"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

// TraderAPI autoorder通用报单接口
type TraderAPI interface {
	Order(d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error)
	FAK(d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error)
	Cancel(localID autoorder.OrderID) error
}

// MockTrader 测试Mock接口
type MockTrader struct {
	orderID int64
}

func (td *MockTrader) mockOrder(name string, d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error) {
	log.Printf("MockTrader.%s called: %s %d@%f\n", name, d.Name(), vol, price)

	if !autoorder.ValidateVolume(vol) {
		return -1, fmt.Errorf("Invalid volume: %d", vol)
	}

	if !autoorder.ValidatePrice(price) {
		return -1, fmt.Errorf("Invalid price: %f", price)
	}

	return autoorder.OrderID(atomic.AddInt64(&td.orderID, 1)), nil
}

// Order 报单接口
func (td *MockTrader) Order(d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error) {

	return td.mockOrder("Order", d, price, vol)
}

// FAK FAK接口
func (td *MockTrader) FAK(d autoorder.Direction, price float64, vol int64) (autoorder.OrderID, error) {
	return td.mockOrder("FAK", d, price, vol)
}

// Cancel 撤单接口
func (td *MockTrader) Cancel(localID autoorder.OrderID) error {
	log.Printf("MockTrader.Cancel called with LocalID: %d\n", localID)
	return nil
}

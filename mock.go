package autoorder

import (
	"fmt"
	"log"
	"sync/atomic"
)

// MockTrader 测试Mock接口
type MockTrader struct {
	orderID int64
}

// Login 登录接口
func (td *MockTrader) Login(loginInfo map[string]string) error {
	return nil
}

// Logout 登出接口
func (td *MockTrader) Logout() error {
	return nil
}

func (td *MockTrader) mockOrder(name string, d Direction, price float64, vol int64) (OrderID, error) {
	log.Printf("MockTrader.%s called: %s %d@%f\n", name, d.Name(), vol, price)

	if !ValidateVolume(vol) {
		return -1, fmt.Errorf("Invalid volume: %d", vol)
	}

	if !ValidatePrice(price) {
		return -1, fmt.Errorf("Invalid price: %f", price)
	}

	return OrderID(atomic.AddInt64(&td.orderID, 1)), nil
}

// Order 报单接口
func (td *MockTrader) Order(d Direction, price float64, vol int64) (OrderID, error) {

	return td.mockOrder("Order", d, price, vol)
}

// FAK FAK接口
func (td *MockTrader) FAK(d Direction, price float64, vol int64) (OrderID, error) {
	return td.mockOrder("FAK", d, price, vol)
}

// Cancel 撤单接口
func (td *MockTrader) Cancel(localID OrderID) error {
	log.Printf("MockTrader.Cancel called with LocalID: %d\n", localID)
	return nil
}

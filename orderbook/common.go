package orderbook

import "fmt"

type direction uint8

const (
	// Buy 买方向
	Buy direction = iota
	// Sell 卖方向
	Sell
)

type priceItem struct {
	Price  float64
	Volume int64
}

type identity struct {
	ExchangeID   string
	InstrumentID string
}

func (id *identity) Identity() string {
	return fmt.Sprintf("%s.%s", id.ExchangeID, id.InstrumentID)
}
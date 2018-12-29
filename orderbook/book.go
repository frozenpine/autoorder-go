package orderbook

import (
	"gitlab.quantdo.cn/yuanyang/autoorder/trader"
)

// Book 订单部
type Book struct {
	identity
	spread
	trader         *trader.TraderAPI
	LastPrice      float64
	TickPrice      float64
	Volume         int64
	MaxVolPerOrder int64
	Asks           *page
	Bids           *page
}

// CreateOrderBook OrderBook工厂函数
func CreateOrderBook(exchangeID, instrumentID string, maxVol int, tick float64, traderAPI *trader.TraderAPI) *Book {
	book := Book{trader: traderAPI, identity: identity{ExchangeID: exchangeID, InstrumentID: instrumentID}}

	book.Asks = createPage(Sell, &book)
	book.Bids = createPage(Buy, &book)

	return &book
}

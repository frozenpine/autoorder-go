package orderbook

import (
	"fmt"
	"log"
	"math"

	"gitlab.quantdo.cn/yuanyang/autoorder"

	"gitlab.quantdo.cn/yuanyang/autoorder/trader"
)

type identity struct {
	ExchangeID   string
	InstrumentID string
}

func (id *identity) Identity() string {
	return fmt.Sprintf("%s.%s", id.ExchangeID, id.InstrumentID)
}

func validateVolume(vol int64) bool {
	return vol != 0
}

func validatePrice(price float64) bool {
	return price != 0 && price != math.MaxFloat64
}

// Book 订单簿
type Book struct {
	identity
	spread
	trader         *trader.TraderAPI
	MaxVolPerOrder int64
	Asks           *page
	Bids           *page
}

// Update 更新订单簿中的价格和量
func (ob *Book) Update(d autoorder.Direction, price float64, volume int64) {
	var dst, oppsite *page
	switch d {
	case autoorder.Sell:
		dst = ob.Asks
		oppsite = ob.Bids
	case autoorder.Buy:
		dst = ob.Bids
		oppsite = ob.Asks
	default:
		panic("Invalid direction.")
	}

	for oppsite.Overlapped(price) {
		lvl := oppsite.PopLevel()
		log.Println(lvl)
		lvl.Remove()
	}

	_, err := dst.GetLevel(price)

	if err != nil {
		dst.ModifyLevel(price, volume)
	} else {
		dst.AddLevel(price, volume)
	}
}

// Rebuild 根据已有委托重建订单簿
func (ob *Book) Rebuild(d autoorder.Direction, price float64, vol int64, localID autoorder.OrderID, sysID int64) {

}

// CreateOrderBook OrderBook工厂函数
func CreateOrderBook(exchangeID, instrumentID string, maxVol int64, tick, open float64, traderAPI *trader.TraderAPI) *Book {
	if !validateVolume(maxVol) || !validatePrice(tick) || !validatePrice(open) {
		return nil
	}

	book := Book{
		trader:   traderAPI,
		identity: identity{ExchangeID: exchangeID, InstrumentID: instrumentID},
		spread:   spread{TickPrice: tick}}

	book.initSpread(open, 0, 0, 0, 0)

	book.Asks = createPage(autoorder.Sell, &book)
	book.Bids = createPage(autoorder.Buy, &book)

	return &book
}

package orderbook

import (
	"fmt"
	"log"
	"math"

	"gitlab.quantdo.cn/yuanyang/autoorder"

	"gitlab.quantdo.cn/yuanyang/autoorder/trader"
)

type identity struct {
	ExchangeID   string `json:"ExchangeID"`
	InstrumentID string `json:"InstrumentID"`
}

func (id *identity) Identity() string {
	return fmt.Sprintf("%s.%s", id.ExchangeID, id.InstrumentID)
}

func validateVolume(vol int64) bool {
	return vol > 0
}

func validatePrice(price float64) bool {
	return price != 0 && price != math.MaxFloat64
}

// Book 订单簿
type Book struct {
	identity
	spread
	trader         trader.TraderAPI
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
	// todo: 重建订单簿逻辑
}

// Snapshot 获取订单簿快照
func (ob *Book) Snapshot() autoorder.Snapshot {
	rtn := ob.spread.Snapshot()

	rtn["ExchangeID"] = ob.ExchangeID
	rtn["InstrumentID"] = ob.InstrumentID

	rtn["Asks"] = ob.Asks.Snapshot()
	rtn["Bids"] = ob.Bids.Snapshot()

	rtn["MaxVolPerOrder"] = ob.MaxVolPerOrder

	return rtn
}

// CreateOrderBook OrderBook工厂函数
func CreateOrderBook(exchangeID, instrumentID string, maxVol int64, tick, open float64, api trader.TraderAPI) *Book {
	if !validateVolume(maxVol) || !validatePrice(tick) || !validatePrice(open) {
		return nil
	}

	book := Book{
		trader:         api,
		identity:       identity{ExchangeID: exchangeID, InstrumentID: instrumentID},
		spread:         spread{TickPrice: tick},
		MaxVolPerOrder: maxVol}

	book.Asks = newPage(autoorder.Sell, maxVol, api)
	book.Bids = newPage(autoorder.Buy, maxVol, api)

	book.initSpread(&book, open, 0, 0, 0, 0)

	return &book
}

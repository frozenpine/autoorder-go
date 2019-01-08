package orderbook

import (
	"fmt"
	"log"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type identity struct {
	ExchangeID   string
	InstrumentID string
}

func (id *identity) Identity() string {
	return fmt.Sprintf("%s.%s", id.ExchangeID, id.InstrumentID)
}

// Book 订单簿
type Book struct {
	identity
	spread
	trader         autoorder.TraderAPI
	MaxVolPerOrder int64
	Asks           *page
	Bids           *page
}

// Update 更新订单簿中的价格和量
func (ob *Book) Update(d autoorder.Direction, price float64, volume int64) {
	if !autoorder.ValidatePrice(price) || !autoorder.ValidateVolume(volume) {
		return
	}

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

	defer ob.UpdateBlock(d, price)

	for oppsite.Overlapped(price) {
		lvl := oppsite.PopLevel()

		log.Println(lvl.Snapshot())
	}

	_, err := dst.GetLevel(price)

	if err != nil {
		dst.ModifyLevel(price, volume)
	} else {
		dst.MakeLevel(price, volume, true)
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
func CreateOrderBook(exchangeID, instrumentID string, maxVol int64, tick, open float64, api autoorder.TraderAPI) *Book {
	if !autoorder.ValidateVolume(maxVol) || !autoorder.ValidatePrice(tick) || !autoorder.ValidatePrice(open) {
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

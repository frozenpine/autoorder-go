package orderbook

import (
	"math"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

const (
	blockVolume int64   = 10000
	floatFact   float64 = 0.2
)

type spread struct {
	MakeBlock     bool
	BlockTick     int
	TickPrice     float64
	OpenPrice     float64
	ClosePrice    float64
	HightestPrice float64
	LowestPrice   float64
	LimitPrice    float64
	StopPrice     float64
	orderBook     *Book
	ceilBlock     *level
	floorBlock    *level
}

func (sp *spread) initSpread(ob *Book, open, high, low, limit, stop float64) {
	sp.OpenPrice = open

	sp.orderBook = ob

	if autoorder.ValidatePrice(limit) {
		sp.LimitPrice = limit
	} else {
		sp.LimitPrice = math.MaxFloat64
	}

	if autoorder.ValidatePrice(stop) {
		sp.StopPrice = stop
	} else {
		sp.StopPrice = 0
	}

	if autoorder.ValidatePrice(high) && high <= sp.LimitPrice {
		sp.HightestPrice = high
	} else {
		sp.HightestPrice = 0
	}

	if autoorder.ValidatePrice(low) && low >= sp.StopPrice {
		sp.LowestPrice = low
	} else {
		sp.LowestPrice = math.MaxFloat64
	}
}

func (sp *spread) calculateCeil(price float64) float64 {
	if sp.BlockTick == 0 {
		sp.BlockTick = 10
	}

	var blockPrice, percentPrice float64 = price + sp.TickPrice*float64(sp.BlockTick), math.MaxFloat64

	if autoorder.ValidatePrice(sp.LimitPrice) {
		tickLevels := (sp.LimitPrice - price) / sp.TickPrice
		percentPrice = autoorder.NormalizePrice(sp.LimitPrice-tickLevels*floatFact*sp.TickPrice, sp.TickPrice, autoorder.RoundUp)
	}

	return math.Max(math.Min(blockPrice, percentPrice), price)
}

func (sp *spread) calculateFloor(price float64) float64 {
	if sp.BlockTick == 0 {
		sp.BlockTick = 10
	}

	var blockPrice, percentPrice float64 = price - sp.TickPrice*float64(sp.BlockTick), 0

	if autoorder.ValidatePrice(sp.StopPrice) {
		tickLevels := (price - sp.StopPrice) / sp.TickPrice
		percentPrice = autoorder.NormalizePrice(sp.StopPrice+tickLevels*floatFact*sp.TickPrice, sp.TickPrice, autoorder.RoundDown)
	}

	return math.Min(math.Max(blockPrice, percentPrice), price)
}

func (sp *spread) UpdateBlock(d autoorder.Direction, price float64) {
	if !sp.MakeBlock {
		return
	}

	var blockPrice float64
	var blockLevel *level
	var checkPrice func() bool
	var funcRenew func()

	switch d {
	case autoorder.Buy:
		blockPrice = sp.calculateFloor(price)

		blockLevel = sp.floorBlock

		checkPrice = func() bool {
			return blockLevel == nil || blockPrice < blockLevel.LevelPrice
		}

		funcRenew = func() {
			lvl := newLevel(blockPrice, blockVolume, sp.orderBook.Bids, false)
			sp.floorBlock = lvl
		}
	case autoorder.Sell:
		blockPrice = sp.calculateCeil(price)

		blockLevel = sp.ceilBlock

		checkPrice = func() bool {
			return blockLevel == nil || blockPrice > blockLevel.LevelPrice
		}

		funcRenew = func() {
			lvl := newLevel(blockPrice, blockVolume, sp.orderBook.Asks, false)
			sp.ceilBlock = lvl
		}
	default:
		panic("Invalid direction.")
	}

	if !autoorder.ValidatePrice(blockPrice) || !checkPrice() {
		return
	}

	funcRenew()

	if blockLevel != nil {
		blockLevel.CancelAll()
	}
}

func (sp *spread) Snapshot() autoorder.Snapshot {
	rtn := autoorder.Snapshot(make(map[string]interface{}))

	rtn["MakeBlock"] = sp.MakeBlock
	rtn["BlockTick"] = sp.BlockTick
	rtn["TickPrice"] = sp.TickPrice
	rtn["OpenPrice"] = sp.OpenPrice
	rtn["ClosePrice"] = sp.ClosePrice
	rtn["HightestPrice"] = sp.HightestPrice
	rtn["LowestPrice"] = sp.LowestPrice
	rtn["LimitPrice"] = sp.LimitPrice
	rtn["StopPrice"] = sp.StopPrice

	rtn["CeiLevel"] = sp.ceilBlock.Snapshot()
	rtn["FloorLevel"] = sp.floorBlock.Snapshot()

	return rtn
}

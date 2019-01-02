package orderbook

import (
	"math"

	"gitlab.quantdo.cn/yuanyang/autoorder"
)

type spread struct {
	makeBlock     bool
	blockTick     int
	TickPrice     float64
	OpenPrice     float64
	ClosePrice    float64
	HightestPrice float64
	LowerestPrice float64
	LimitPrice    float64
	StopPrice     float64
	ceilBlock     *level
	floorBlock    *level
}

func (sp *spread) initSpread(open, high, low, limit, stop float64) {
	sp.OpenPrice = open

	if validatePrice(limit) {
		sp.LimitPrice = limit
	} else {
		sp.LimitPrice = math.MaxFloat64
	}

	if validatePrice(stop) {
		sp.StopPrice = stop
	} else {
		sp.StopPrice = 0
	}

	if validatePrice(high) && high <= sp.LimitPrice {
		sp.HightestPrice = high
	} else {
		sp.HightestPrice = 0
	}

	if validatePrice(low) && low >= sp.StopPrice {
		sp.LowerestPrice = low
	} else {
		sp.LowerestPrice = math.MaxFloat64
	}
}

func (sp *spread) calculateCeil(price float64) float64 {
	return math.MaxFloat64
}

func (sp *spread) calculateFloor(price float64) float64 {
	return 0
}

func (sp *spread) UpdateBlock(d autoorder.Direction, price float64) {
	if !sp.makeBlock {
		return
	}

	switch d {
	case autoorder.Buy:
		blockPrice := sp.calculateFloor(price)
		if sp.floorBlock == nil || blockPrice < sp.floorBlock.LevelPrice {
			sp.floorBlock.Remove()

		}
	case autoorder.Sell:
	default:
		panic("Invalid direction.")
	}
}

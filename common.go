package autoorder

import "math"

// OrderID 本地报单编号
type OrderID int64

// Direction 买卖方向
type Direction uint8

const (
	// Buy 买方向
	Buy Direction = iota
	// Sell 卖方向
	Sell
)

// Opposite 获取买卖方向的反方向
func (d Direction) Opposite() Direction {
	switch d {
	case Buy:
		return Sell
	case Sell:
		return Buy
	default:
		panic("Invalid direction.")
	}
}

// Name 获取买卖方向的名字
func (d Direction) Name() string {
	switch d {
	case Buy:
		return "Buy"
	case Sell:
		return "Sell"
	default:
		panic("Invalid direction.")
	}
}

// Tick 行情Tick
type Tick struct {
	Symbol        string  `json:"Symbol"`
	AskPrice      float64 `json:"AskPrice"`
	AskVolume     int64   `json:"AskVolume"`
	BidPrice      float64 `json:"BidPrice"`
	BidVolume     int64   `json:"BidVolume"`
	LastPrice     float64 `json:"LastPrice"`
	LastVolume    int64   `json:"LastVolume"`
	HightestPrice float64 `json:"HightestPrice"`
	LowestPrice   float64 `json:"LowestPrice"`
	MTS           int     `json:"MTS"`
}

// Order 逐笔委托更新
type Order struct {
	Price  float64 `json:"Price"`
	Volume int64   `json:"Volume"`
	Count  int     `json:"Count"`
}

// Candle K线数据
type Candle struct {
	Open   float64 `json:"Open"`
	Close  float64 `json:"Close"`
	High   float64 `json:"High"`
	Low    float64 `json:"Low"`
	Volume int64   `json:"Volume"`
	MTS    int     `json:"MTS"`
}

// Trade 成交数据
type Trade struct {
	TradeID int64   `json:"TradeID"`
	Price   float64 `json:"Price"`
	Volume  int64   `json:"Volume"`
	MTS     int     `json:"MTS"`
}

// MaxFloat64 查找一组float64中的最大值
func MaxFloat64(f ...float64) float64 {
	var max float64
	max = 0
	for _, v := range f {
		max = math.Max(max, v)
	}
	return max
}

// MinFloat64 查找一组float64中的最小值
func MinFloat64(f ...float64) float64 {
	var min float64
	min = math.MaxFloat64
	for _, v := range f {
		min = math.Min(min, v)
	}
	return min
}

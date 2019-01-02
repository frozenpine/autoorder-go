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

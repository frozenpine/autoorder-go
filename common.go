package autoorder

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

func (d Direction) opposite() Direction {
	switch d {
	case Buy:
		return Sell
	case Sell:
		return Buy
	default:
		panic("Invalid direction.")
	}
}

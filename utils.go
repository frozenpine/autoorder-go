package autoorder

import "math"

// Snapshot 数据快照
type Snapshot map[string]interface{}

// SnapshotMerge 合并两个快照至dst, 如存在Key重复, src的Value将覆盖dst的Value
func SnapshotMerge(dst Snapshot, src Snapshot) Snapshot {
	for key, value := range src {
		dst[key] = value
	}

	return dst
}

// ValidateVolume 校验Volume合法性, <=0 非法
func ValidateVolume(vol int64) bool {
	return vol > 0
}

// ValidatePrice 校验价格合法性, 0 | MaxFloat64 非法
func ValidatePrice(price float64) bool {
	return price != 0 && price != math.MaxFloat64
}

type roundMode uint8

const (
	// RoundDefault 默认四舍五入
	RoundDefault roundMode = iota
	// RoundUp 向上取整
	RoundUp
	// RoundDown 向下取整
	RoundDown
)

// NormalizePrice 将价格以指定的方式取整至tickPrice的整数倍
func NormalizePrice(price, tickPrice float64, round roundMode) float64 {
	multiple := price / tickPrice

	switch round {
	case RoundDefault:
		multiple = math.Round(multiple)
	case RoundUp:
		multiple = math.Ceil(multiple)
	case RoundDown:
		multiple = math.Floor(multiple)
	default:
		multiple = math.Round(multiple)
	}

	return tickPrice * multiple
}

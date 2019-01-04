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

func ValidateVolume(vol int64) bool {
	return vol > 0
}

func ValidatePrice(price float64) bool {
	return price != 0 && price != math.MaxFloat64
}

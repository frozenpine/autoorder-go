package autoorder

// Snapshot 数据快照
type Snapshot map[string]interface{}

// SnapshotMerge 合并两个快照至dst, 如存在Key重复, src的Value将覆盖dst的Value
func SnapshotMerge(dst Snapshot, src Snapshot) Snapshot {
	for key, value := range src {
		dst[key] = value
	}

	return dst
}

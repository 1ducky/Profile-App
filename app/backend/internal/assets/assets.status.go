package assets

type AssetStatus string

const (
	AssetStatusTemp    AssetStatus = "temp"
	AssetStatusPending AssetStatus = "pending"
	AssetStatusActive  AssetStatus = "active"
	AssetStatusOrphan  AssetStatus = "orphan"
	AssetStatusDeleted AssetStatus = "deleted"
	AssetStatusFailed  AssetStatus = "failed"
)

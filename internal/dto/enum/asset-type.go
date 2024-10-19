package enum

// AssetType represents the type of asset
type AssetType string

// Asset types
const (
	AssetTypeVehicle AssetType = "vehicle"
	AssetTypeHouse   AssetType = "house"
	AssetTypeCrypto  AssetType = "crypto"
	AssetTypeOther   AssetType = "other"
)

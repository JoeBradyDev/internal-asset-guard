package dto

// Asset Class

type CreateAssetClass struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

type UpdateAssetClass struct {
	Name       *string `json:"name"`
	Definition *string `json:"definition"`
}

type AssetClass struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

// Criticality

type CreateCriticality struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type UpdateCriticality struct {
	Name  *string `json:"name"`
	Value *int    `json:"value"`
}

type Criticality struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Device Type

type CreateDeviceType struct {
	AssetClassID int    `json:"asset_class_id"`
	Name         string `json:"name"`
}

type UpdateDeviceType struct {
	AssetClassID *int    `json:"asset_class_id"`
	Name         *string `json:"name"`
}

type DeviceType struct {
	ID           int    `json:"id"`
	AssetClassID int    `json:"asset_class_id"`
	Name         string `json:"name"`
}

// Asset Source

type CreateAssetSource struct {
	Name string `json:"name"`
}

type UpdateAssetSource struct {
	Name *string `json:"name"`
}

type AssetSource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


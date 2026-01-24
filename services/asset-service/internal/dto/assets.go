package dtos

import "time"

// --- CORE ASSET ---

type CreateAsset struct {
	Name          string          `json:"name"`
	AssetClassID  int             `json:"asset_class_id"`
	CriticalityID int             `json:"criticality_id"`
	DeviceInfo    *DeviceDetail   `json:"device_info,omitempty"`
	NetworkInfo   *NetworkDetail  `json:"network_info,omitempty"`
	SoftwareInfo  *SoftwareDetail `json:"software_info,omitempty"`
}

type UpdateAsset struct {
	Name          *string         `json:"name"`
	AssetClassID  *int            `json:"asset_class_id"`
	CriticalityID *int            `json:"criticality_id"`
	DeviceInfo    *DeviceDetail   `json:"device_info,omitempty"`
	NetworkInfo   *NetworkDetail  `json:"network_info,omitempty"`
	SoftwareInfo  *SoftwareDetail `json:"software_info,omitempty"`
}

type Asset struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	AssetClassID    int             `json:"asset_class_id"`
	AssetClassLabel string          `json:"asset_class_label,omitempty"`
	CriticalityID   int             `json:"criticality_id"`
	CriticalityName string          `json:"criticality_label,omitempty"`
	CreatedAt       time.Time       `json:"created_at"`
	DeviceInfo      *DeviceDetail   `json:"device_info,omitempty"`
	NetworkInfo     *NetworkDetail  `json:"network_info,omitempty"`
	SoftwareInfo    *SoftwareDetail `json:"software_info,omitempty"`
	TotalIssues     int             `json:"total_issues,omitempty"`
}

// --- POLYMORPHIC DETAILS ---

type DeviceDetail struct {
	Hostname     *string    `json:"hostname"`
	DeviceTypeID *int       `json:"device_type_id"`
	DeviceType   string     `json:"device_type_label,omitempty"`
	IPAddress    *string    `json:"ip_address"`
	MacAddress   *string    `json:"mac_address"`
	OSName       *string    `json:"os_name"`
	OSVersion    *string    `json:"os_version"`
	HardwareCPE  *string    `json:"hardware_cpe"`
	LastSeen     *time.Time `json:"last_seen"`
}

type NetworkDetail struct {
	ManagementIP    *string    `json:"management_ip"`
	DeviceTypeID    *int       `json:"device_type_id"`
	DeviceType      string     `json:"device_type_label,omitempty"`
	MacAddress      *string    `json:"mac_address"`
	FirmwareVersion *string    `json:"firmware_version"`
	ModelNumber     *string    `json:"model_number"`
	SerialNumber    *string    `json:"serial_number"`
	TotalPorts      *int       `json:"total_ports"`
	LastSeen        *time.Time `json:"last_seen"`
}

type SoftwareDetail struct {
	Name        *string    `json:"name"`
	OSName      *string    `json:"os_name"`
	OSVersion   *string    `json:"os_version"`
	Version     *string    `json:"version"`
	Vendor      *string    `json:"vendor"`
	SoftwareCPE *string    `json:"software_cpe"`
	LastSeen    *time.Time `json:"last_seen"`
}

package descriptionalSpec

// KS X 3286 노드/디바이스 규격 구조체 목록

// CommSpec
type CommSpec map[string]ProtocolDetail

type CommonField struct {
	Class          string      `json:"Class"`
	Model          string      `json:"Model,omitempty"`
	Name           string      `json:"Name,omitempty"`
	DeviceType     string      `json:"Type"`
	ValueRange     *ValueRange `json:"ValueRange,omitempty"`
	CommSpec       CommSpec    `json:"CommSpec"`
	VendorSpecific string      `json:"VendorSpecific,omitempty"`
}

type ProtocolDetail struct {
	Read  *RegisterInfo `json:"read,omitempty"`
	Write *RegisterInfo `json:"write,omitempty"`
}

type ValueRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type RegisterInfo struct {
	StartingRegister int      `json:"starting-register"`
	Items            []string `json:"items"`
}

type NodeSpec struct {
	CommonField
	Devices []string `json:"Devices,omitempty"`
}

type SensorSpec struct {
	CommonField
	SignificantDigit *int   `json:"SignificantDigit,omitempty"`
	ValueType        string `json:"ValueType"`
	Channel          *int   `json:"Channel,omitempty"`
}

type ActuatorSpec struct {
	CommonField
	ValueUnit string `json:"ValueUnit"`
	Channel   *int   `json:"Channel,omitempty"`
}

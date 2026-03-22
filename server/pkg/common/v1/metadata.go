package commonv1

// Metadata tương ứng với proto Metadata
type Metadata struct {
	RequestID  string `json:"request_id,omitempty"`
	TraceID    string `json:"trace_id,omitempty"`
	ServerTime int64  `json:"server_time,omitempty"` // Unix timestamp hoặc milliseconds
	Locale     string `json:"locale,omitempty"`

	// Optional advanced fields
	Region   string `json:"region,omitempty"`   // ap-sg, ap-vn
	Degraded bool   `json:"degraded,omitempty"` // hệ thống đang degraded
}

package types

// Header ...
type Header struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
}

// HealthzResponse ...
type HealthzResponse struct {
	Result bool `json:"result"`
}

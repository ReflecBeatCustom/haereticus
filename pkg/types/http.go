package types

// Header ...
type Header struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
}

// HealthzResponse ...
type HealthzResponse struct {
	Result bool `json:"result"`
}

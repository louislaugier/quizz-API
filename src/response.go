package response

// Response export
type Response struct {
	Limit    string        `json:"limit,omitempty"`
	Username string        `json:"username,omitempty"`
	Error    error         `json:"error,omitempty"`
	Data     []interface{} `json:"data"`
}

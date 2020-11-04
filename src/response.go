package response

// Response export
type Response struct {
	Limit     string        `json:"limit,omitempty"`
	Username  string        `json:"username,omitempty"`
	Error     error         `json:"error,omitempty"`
	Questions []interface{} `json:"questions,omitempty"`
	Scores    []interface{} `json:"scores,omitempty"`
}

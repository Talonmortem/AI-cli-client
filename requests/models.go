package requests

type Request struct {
	Prompt string `json:"prompt"`
	Text   string `json:"transcript,omitempty"`

	Role     string `json:"role"`
	Model    string `json:"model"` // e.g., "gpt-4"
	Provider string `json:"provider"`
	APIKey   string `json:"api_key,omitempty"`
}

type Response struct {
	Text string `json:"text"`
}

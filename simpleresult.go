package goresult

type SimpleResult struct {
	Success  bool      `json:"success"`
	Messages []Message `json:"messages,omitempty"`
}

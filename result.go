package goresult

type Result[T any] struct {
	Success  bool      `json:"success"`
	Data     T         `json:"data,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}

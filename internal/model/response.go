package model

type Response[Type any] struct {
	Status int    `json:"status"`
	Data   Type   `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

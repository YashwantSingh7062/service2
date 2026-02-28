package model

type Order struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}

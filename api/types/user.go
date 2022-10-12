package types

type User struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Key    string `json:"key,omitempty"`
	Secret string `json:"secret,omitempty"`
}
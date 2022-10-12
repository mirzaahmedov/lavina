package types

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}
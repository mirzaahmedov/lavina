package types

type Response struct {
	Data    interface{} `json:"data"`
	IsOk    bool        `json:"isOk"`
	Message string      `json:"message"`
}
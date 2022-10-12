package types

type Data interface {
	User | Book
}

type Response[T Data] struct {
	Data    *T     `json:"data"`
	IsOk    bool   `json:"isOk"`
	Message string `json:"message"`
}
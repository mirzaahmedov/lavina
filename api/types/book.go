package types

const (
	StatusNew = iota
	StatusReading
	StatusFinished
)

type BookInfo struct {
	Id        string `json:"id,omitempty"`
	Isbn      string `json:"isbn,omitempty"`
	Title     string `json:"title,omitempty"`
	Author    string `json:"author,omitempty"`
	Published string `json:"published,omitempty"`
	Pages     int    `json:"pages,omitempty"`
}

type Book struct {
	Book   BookInfo `json:"book,omitempty"`
	UserId int      `json:"user_id,omitempty"`
	Status int      `json:"status"`
}
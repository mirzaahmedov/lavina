package types

const (
	StatusNew = iota
	StatusReading
	StatusFinished
)

type Book struct {
	Book struct {
		Id        string `json:"id"`
		Isbn      string `json:"isbn"`
		Title     string `json:"title"`
		Author    string `json:"author"`
		Published int    `json:"published"`
		Pages     int    `json:"pages"`
	}
	Status int `json:"status"`
}
package test_app

type Message struct {
	Date   string `json:"date" binding:"required"`
	Number string `json:"number" binding:"required"`
	Speed  string `json:"speed" binding:"required"`
}

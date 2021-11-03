package test_app

type OverSpeedQuery struct {
	Date  string `json:"date" binding:"required"`
	Speed string `json:"speed" binding:"required"`
}

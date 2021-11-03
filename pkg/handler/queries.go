package handler

import (
	"net/http"
	"time"

	test_app "github.com/AnnZh/test-app"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type getCars struct {
	Data []test_app.Message `json:"data"`
}

func (h *Handler) getOverspeedCars(c *gin.Context) {
	var input test_app.OverSpeedQuery

	if !CheckTime() {
		newErrorResponse(c, http.StatusForbidden, "not available at this time")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cars, err := h.services.Queries.GetOverspeedCars(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCars{
		Data: cars,
	})
}

func (h *Handler) getMinMaxSpeedCars(c *gin.Context) {
	var input test_app.MinMaxQuery

	if !CheckTime() {
		newErrorResponse(c, http.StatusForbidden, "not available at this time")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cars, err := h.services.Queries.GetMinMaxSpeedCars(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCars{
		Data: cars,
	})
}

func CheckTime() bool {
	today := time.Now()
	stringTodayTime := today.Format("15:04")
	todayTime, _ := time.Parse("15:04", stringTodayTime)

	startTime, _ := time.Parse("15:04", viper.GetString("time.start"))

	endTime, _ := time.Parse("15:04", viper.GetString("time.end"))

	if !(todayTime.After(startTime) && todayTime.Before(endTime)) {
		return false
	} else {
		return true
	}
}

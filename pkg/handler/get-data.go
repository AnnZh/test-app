package handler

import (
	"net/http"

	test_app "github.com/AnnZh/test-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getData(c *gin.Context) {
	var input test_app.Message

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Data.GetData(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}

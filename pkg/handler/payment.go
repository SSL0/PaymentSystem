package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResponseBody struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float32 `json:"amount"`
}

func (h *Handler) Send(c *gin.Context) {
	body := ResponseBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input data")
		return
	}

	if body.From == "" || body.To == "" || body.Amount <= 0.0 {
		c.JSON(http.StatusBadRequest, "invalid input data")
		return
	}

	err := h.services.Send(body.From, body.To, body.Amount)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "internal Server Error")
		return
	}

	c.JSON(http.StatusOK, "successful")
}

func (h *Handler) GetBalance(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, "invalid input data")
		return
	}
	balance, err := h.services.GetBalance(address)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"balance": balance,
	})
}

func (h *Handler) GetLast(c *gin.Context) {
	N := c.Query("count")
	count, err := strconv.Atoi(N)
	if err != nil {
		c.JSON(http.StatusBadRequest, "count must be integer")
		return
	}

	if count <= 0 {
		c.JSON(http.StatusBadRequest, "invalid input data")
		return
	}

	transactions, err := h.services.GetLast(count)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": transactions,
	})
}

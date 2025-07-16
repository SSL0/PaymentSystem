package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SendRequestBody struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float32 `json:"amount"`
}

// Send godoc
// @Summary     Send handler function
// @Description makes transaction to send some amount of money from one wallet to other
// @Tags 		transactions
// @Accept 		json
// @Produce 	json
// @Param 		request body SendRequestBody true "Transaction details"
// @Success 	200 {object} map[string]interface{} "Returns success message"
// @Failure		400 {object} map[string]string "Invalid input data"
// @Failure 	500 {object} map[string]string "Internal server error"
// @Router 		/send [post]
func (h *Handler) Send(c *gin.Context) {
	body := SendRequestBody{}

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

// GetBalance godoc
// @Summary     Get balance by wallet address
// @Description returns the amount of funds stored in the wallet
// @Tags		wallet
// @Accept 		json
// @Produce 	json
// @Param 		address path string true "Wallet address"
// @Success 	200 {object} map[string]interface{} "Returns wallet address and balance"
// @Failure 	400 {object} map[string]string "Invalid wallet address (empty)"
// @Failure 	500 {object} map[string]string "Internal server error"
// @Router 		/wallet/{address}/balance [get]
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

// GetBalance godoc
// @Summary      Get last N transactions
// @Description  returns information about the N most recent transfers of funds
// @Tags		 transactions
// @Accept       json
// @Produce      json
// @Param        count   query     int  true  "Number of transactions to fetch (must be positive)"
// @Success      200     {object}  map[string]interface{}  "Returns list of transactions"
// @Failure      400     {object}  map[string]string       "Invalid count (not an integer or <= 0)"
// @Failure      500     {object}  map[string]string       "Internal server error"
// @Router       /transactions [get]
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

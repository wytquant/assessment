package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wytquant/assessment/helpers"
	"github.com/wytquant/assessment/requests"
	"github.com/wytquant/assessment/src/expense/services"
)

type expenseHandler struct {
	expenseService services.ExpenseService
}

func NewExpenseHandler(expenseService services.ExpenseService) expenseHandler {
	return expenseHandler{expenseService: expenseService}
}

func (h expenseHandler) CreateExpense(c *gin.Context) {
	var expense requests.ExpenseRequest
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	expsResponse, err := h.expenseService.CreateExpense(expense)
	if err != nil {
		appErr, ok := err.(*helpers.AppError)
		if ok {
			c.JSON(appErr.StatusCode, appErr.Message)
		}
		return
	}

	c.JSON(http.StatusCreated, expsResponse)
}

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
			c.JSON(appErr.StatusCode, gin.H{"message": appErr.Message})
		}
		return
	}

	c.JSON(http.StatusCreated, expsResponse)
}

func (h expenseHandler) GetExpenseByID(c *gin.Context) {
	expenseResp, err := h.expenseService.GetExpenseByID(c.Param("id"))
	if err != nil {
		appErr, ok := err.(*helpers.AppError)
		if ok {
			c.JSON(appErr.StatusCode, gin.H{"message": appErr.Message})
		}
		return
	}

	c.JSON(http.StatusOK, expenseResp)
}

func (h expenseHandler) UpdateExpenseByID(c *gin.Context) {
	var expenseReq requests.ExpenseRequest
	if err := c.ShouldBindJSON(&expenseReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	expenseResp, err := h.expenseService.UpdateExpenseByID(c.Param("id"), expenseReq)
	if err != nil {
		appErr, ok := err.(*helpers.AppError)
		if ok {
			c.JSON(appErr.StatusCode, gin.H{"message": appErr.Message})
		}
		return
	}

	c.JSON(http.StatusOK, expenseResp)
}

func (h expenseHandler) GetAllExpenses(c *gin.Context) {
	expenseResp, err := h.expenseService.GetExpenses()
	if err != nil {
		appErr, ok := err.(*helpers.AppError)
		if ok {
			c.JSON(appErr.StatusCode, gin.H{"message": appErr.Message})
		}
		return
	}

	c.JSON(http.StatusOK, expenseResp)
}

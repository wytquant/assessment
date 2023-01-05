package services

import (
	"github.com/wytquant/assessment/requests"
	"github.com/wytquant/assessment/responses"
)

type ExpenseService interface {
	CreateExpense(requests.ExpenseRequest) (responses.ExpenseResponse, error)
	GetExpenseByID(id string) (responses.ExpenseResponse, error)
	UpdateExpenseByID(id string, expensReq requests.ExpenseRequest) (responses.ExpenseResponse, error)
	GetExpenses() ([]responses.ExpenseResponse, error)
}

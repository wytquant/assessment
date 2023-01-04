package services

import (
	"github.com/wytquant/assessment/requests"
	"github.com/wytquant/assessment/responses"
)

type ExpenseService interface {
	CreateExpense(requests.ExpenseRequest) (*responses.ExpenseResponse, error)
}

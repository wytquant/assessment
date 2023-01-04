package services

import (
	"github.com/stretchr/testify/mock"
	"github.com/wytquant/assessment/requests"
	"github.com/wytquant/assessment/responses"
)

type expenseServiceMock struct {
	mock.Mock
}

func NewExpenseServiceMock() *expenseServiceMock {
	return &expenseServiceMock{}
}

func (s *expenseServiceMock) CreateExpense(expenseReq requests.ExpenseRequest) (*responses.ExpenseResponse, error) {
	args := s.Called()
	return args.Get(0).(*responses.ExpenseResponse), args.Error(1)
}

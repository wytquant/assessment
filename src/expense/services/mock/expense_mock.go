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

func (m *expenseServiceMock) CreateExpense(expenseReq requests.ExpenseRequest) (responses.ExpenseResponse, error) {
	args := m.Called()
	return args.Get(0).(responses.ExpenseResponse), args.Error(1)
}

func (m *expenseServiceMock) GetExpenseByID(id string) (responses.ExpenseResponse, error) {
	args := m.Called(id)
	return args.Get(0).(responses.ExpenseResponse), args.Error(1)
}

func (m *expenseServiceMock) UpdateExpenseByID(id string, expensReq requests.ExpenseRequest) (responses.ExpenseResponse, error) {
	args := m.Called(id, expensReq)
	return args.Get(0).(responses.ExpenseResponse), args.Error(1)
}

func (m *expenseServiceMock) GetExpenses() ([]responses.ExpenseResponse, error) {
	args := m.Called()
	return args.Get(0).([]responses.ExpenseResponse), args.Error(1)
}

package repositories

import (
	"github.com/stretchr/testify/mock"
	"github.com/wytquant/assessment/models"
)

type expenseRepositoryMock struct {
	mock.Mock
}

func NewExpenseReporitoryMock() *expenseRepositoryMock {
	return &expenseRepositoryMock{}
}

func (r *expenseRepositoryMock) Create(expense *models.Expense) error {
	args := r.Called()
	return args.Error(0)
}

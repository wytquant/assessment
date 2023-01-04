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

func (m *expenseRepositoryMock) Create(expense *models.Expense) error {
	args := m.Called()
	return args.Error(0)
}

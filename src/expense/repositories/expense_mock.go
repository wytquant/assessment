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

func (m *expenseRepositoryMock) GetByID(id string) (*models.Expense, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Expense), args.Error(1)
}

func (m *expenseRepositoryMock) UpdateByID(id string, expense models.Expense) (*models.Expense, error) {
	args := m.Called(id, expense)
	return args.Get(0).(*models.Expense), args.Error(1)
}

func (r expenseRepositoryMock) GetAll() (*[]models.Expense, error) {
	return nil, nil
}

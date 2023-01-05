package repositories

import "github.com/wytquant/assessment/models"

type ExpenseRepository interface {
	Create(*models.Expense) error
	GetByID(id string) (*models.Expense, error)
	UpdateByID(id string, expense models.Expense) (*models.Expense, error)
}

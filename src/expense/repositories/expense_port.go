package repositories

import "github.com/wytquant/assessment/models"

type ExpenseRepository interface {
	Create(*models.Expense) error
	GetById(id string) (*models.Expense, error)
}

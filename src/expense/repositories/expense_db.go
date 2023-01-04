package repositories

import (
	"github.com/wytquant/assessment/models"
	"gorm.io/gorm"
)

type expenseRepositoryDB struct {
	db *gorm.DB
}

func NewExpenseRepositoryDB(db *gorm.DB) ExpenseRepository {
	return expenseRepositoryDB{db: db}
}

func (r expenseRepositoryDB) Create(expense *models.Expense) error {
	query := r.db
	if err := query.Create(expense).Error; err != nil {
		return err
	}

	return nil
}

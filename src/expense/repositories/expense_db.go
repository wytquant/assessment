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

func (r expenseRepositoryDB) GetByID(id string) (models.Expense, error) {
	var expense models.Expense
	query := r.db
	if err := query.Where("id = $1", id).First(&expense).Error; err != nil {
		return models.Expense{}, err
	}

	return expense, nil
}

func (r expenseRepositoryDB) UpdateByID(id string, expense models.Expense) (models.Expense, error) {
	query := r.db
	expenseDB, err := r.GetByID(id)
	if err != nil {
		return models.Expense{}, err
	}

	if err := query.Model(&expenseDB).Updates(expense).Error; err != nil {
		return models.Expense{}, err
	}

	return expenseDB, nil
}

func (r expenseRepositoryDB) GetAll() ([]models.Expense, error) {
	query := r.db
	var expenses []models.Expense

	if err := query.Find(&expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

package services

import (
	"github.com/jinzhu/copier"
	"github.com/wytquant/assessment/helpers"
	"github.com/wytquant/assessment/models"
	"github.com/wytquant/assessment/requests"
	"github.com/wytquant/assessment/responses"
	"github.com/wytquant/assessment/src/expense/repositories"
)

type expenseService struct {
	expenseRepo repositories.ExpenseRepository
}

func NewExpenseService(expenseRepo repositories.ExpenseRepository) ExpenseService {
	return expenseService{expenseRepo: expenseRepo}
}

func (s expenseService) CreateExpense(expenseReq requests.ExpenseRequest) (*responses.ExpenseResponse, error) {
	var expense models.Expense
	var expenseResp responses.ExpenseResponse

	copier.Copy(&expense, &expenseReq)

	if err := s.expenseRepo.Create(&expense); err != nil {
		return nil, helpers.NewInternalServerError()
	}

	copier.Copy(&expenseResp, &expense)

	return &expenseResp, nil
}

func (s expenseService) GetExpenseByID(id string) (*responses.ExpenseResponse, error) {
	var expenseResp responses.ExpenseResponse

	expense, err := s.expenseRepo.GetByID(id)
	if err != nil {
		return nil, helpers.NewNotFoundError()
	}

	copier.Copy(&expenseResp, &expense)

	return &expenseResp, nil
}

func (s expenseService) UpdateExpenseByID(id string, expensReq requests.ExpenseRequest) (*responses.ExpenseResponse, error) {
	var expense models.Expense
	var expenseResp responses.ExpenseResponse

	copier.Copy(&expense, &expensReq)

	updatedExpense, err := s.expenseRepo.UpdateByID(id, expense)
	if err != nil {
		return nil, helpers.NewNotFoundError()
	}

	copier.Copy(&expenseResp, &updatedExpense)

	return &expenseResp, nil
}

func (s expenseService) GetExpenses() (*[]responses.ExpenseResponse, error) {
	expensesResp := []responses.ExpenseResponse{}

	expenses, err := s.expenseRepo.GetAll()
	if err != nil {
		return nil, helpers.NewInternalServerError()
	}

	copier.Copy(&expensesResp, &expenses)

	return &expensesResp, nil
}

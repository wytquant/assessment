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

func (s expenseService) GetExpenseById(id string) (*responses.ExpenseResponse, error) {
	var expenseResp responses.ExpenseResponse

	expense, err := s.expenseRepo.GetById(id)
	if err != nil {
		return nil, helpers.NewInternalServerError()
	}

	copier.Copy(&expenseResp, &expense)

	return &expenseResp, nil
}

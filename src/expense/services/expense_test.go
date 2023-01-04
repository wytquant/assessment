//go:build unit

package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wytquant/assessment/helpers"
	"github.com/wytquant/assessment/requests"
	"github.com/wytquant/assessment/src/expense/repositories"
	"github.com/wytquant/assessment/src/expense/services"
)

func TestCreateExpense(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		//Arrange
		expenseRepo := repositories.NewExpenseReporitoryMock()
		expenseRepo.On("Create").Return(nil)

		expenseService := services.NewExpenseService(expenseRepo)

		//act
		_, err := expenseService.CreateExpense(requests.ExpenseRequest{})

		//assert
		assert.NoError(t, err)
	})

	t.Run("fail case", func(t *testing.T) {
		//Arrange
		expenseRepo := repositories.NewExpenseReporitoryMock()
		expenseRepo.On("Create").Return(helpers.NewInternalServerError())

		expenseService := services.NewExpenseService(expenseRepo)

		//act
		_, err := expenseService.CreateExpense(requests.ExpenseRequest{})

		//assert
		assert.EqualError(t, err, helpers.NewInternalServerError().Error())
	})
}

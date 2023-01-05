//go:build unit

package services_test

import (
	"net/http"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/wytquant/assessment/helpers"
	"github.com/wytquant/assessment/models"
	"github.com/wytquant/assessment/requests"
	"github.com/wytquant/assessment/responses"
	"github.com/wytquant/assessment/src/expense/repositories"
	"github.com/wytquant/assessment/src/expense/services"
)

func TestCreateExpenseService(t *testing.T) {
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

func TestGetExpenseByIDService(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		//arrange
		id := "1"
		expenseRepo := repositories.NewExpenseReporitoryMock()
		expenseRepo.On("GetByID", id).Return(models.Expense{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}, nil)
		expenseService := services.NewExpenseService(expenseRepo)

		want := responses.ExpenseResponse{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}

		got, err := expenseService.GetExpenseByID(id)

		assert.NoError(t, err)
		if !assert.ObjectsAreEqual(want, got) {
			t.Errorf("not equal. want: %#v, got: %#v", want, got)
		}
	})

	t.Run("fail case due to record not found", func(t *testing.T) {
		//arrange
		id := "1"
		expenseRepo := repositories.NewExpenseReporitoryMock()
		expenseRepo.On("GetByID", id).Return(models.Expense{}, helpers.NewNotFoundError())

		expenseService := services.NewExpenseService(expenseRepo)

		//act
		_, err := expenseService.GetExpenseByID(id)

		//assert
		appErr, ok := err.(*helpers.AppError)
		if ok {
			assert.Equal(t, http.StatusNotFound, appErr.StatusCode)
		}
		assert.EqualError(t, err, helpers.NewNotFoundError().Error())

	})
}

func TestUpdateExpenseByIDService(t *testing.T) {
	t.Run("update success case", func(t *testing.T) {
		//arrange
		id := "1"
		updatedExpense := models.Expense{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}

		expenseReq := requests.ExpenseRequest{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}

		expenseRepo := repositories.NewExpenseReporitoryMock()
		expenseRepo.On("UpdateByID", id, updatedExpense).Return(models.Expense{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}, nil)

		expenseService := services.NewExpenseService(expenseRepo)

		want := responses.ExpenseResponse{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}

		//act
		got, err := expenseService.UpdateExpenseByID(id, expenseReq)

		//assert
		assert.NoError(t, err)
		if !assert.ObjectsAreEqual(want, got) {
			t.Errorf("not equal. want: %#v, got: %#v", want, got)
		}
	})

	t.Run("fail case bacause expense was not found", func(t *testing.T) {
		//arrange
		id := "1"
		updatedExpense := models.Expense{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}

		expenseReq := requests.ExpenseRequest{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}
		expenseRepo := repositories.NewExpenseReporitoryMock()
		expenseRepo.On("UpdateByID", id, updatedExpense).Return(models.Expense{}, helpers.NewNotFoundError())

		expenseService := services.NewExpenseService(expenseRepo)

		//act
		_, err := expenseService.UpdateExpenseByID(id, expenseReq)

		//assert
		appErr, ok := err.(*helpers.AppError)
		if ok {
			assert.Equal(t, http.StatusNotFound, appErr.StatusCode)
		}
		assert.EqualError(t, err, helpers.NewNotFoundError().Error())
	})
}

func TestGetAllExpensesService(t *testing.T) {
	t.Run("get all success case", func(t *testing.T) {
		//Arrange
		expenseRepo := repositories.NewExpenseReporitoryMock()
		expenseRepo.On("GetAll").Return([]models.Expense{
			{
				ID:     1,
				Title:  "strawberry smoothie",
				Amount: 79,
				Note:   "night market promotion discount 10 bath",
				Tags:   pq.StringArray{"food", "beverage"},
			},
			{
				ID:     2,
				Title:  "strawberry smoothie",
				Amount: 79,
				Note:   "night market promotion discount 10 bath",
				Tags:   pq.StringArray{"food", "beverage"},
			},
		}, nil)

		expenseService := services.NewExpenseService(expenseRepo)

		//act
		got, err := expenseService.GetExpenses()

		//assert
		assert.NoError(t, err)
		assert.NotZero(t, len(got))
	})

	t.Run("get all fail case because internal server error", func(t *testing.T) {
		//Arrange
		expenseRepo := repositories.NewExpenseReporitoryMock()
		expenseRepo.On("GetAll").Return([]models.Expense{}, helpers.NewInternalServerError())

		expenseService := services.NewExpenseService(expenseRepo)

		//act
		got, err := expenseService.GetExpenses()

		//assert
		assert.EqualError(t, err, helpers.NewInternalServerError().Error())
		assert.Equal(t, 0, len(got))
	})
}

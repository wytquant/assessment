//go:build unit

package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/wytquant/assessment/helpers"
	"github.com/wytquant/assessment/responses"
	"github.com/wytquant/assessment/src/expense/handlers"
	services "github.com/wytquant/assessment/src/expense/services/mock"
)

func TestCreateExpenseHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		want := responses.ExpenseResponse{
			ID:     1,
			Title:  "firerice",
			Amount: 100,
			Note:   "new dish",
			Tags:   pq.StringArray{"food"},
		}

		expenseService := services.NewExpenseServiceMock()
		expenseService.On("CreateExpense").Return(&want, nil)

		expenseHandler := handlers.NewExpenseHandler(expenseService)

		r := gin.Default()
		r.POST("/expenses", expenseHandler.CreateExpense)

		payload := strings.NewReader(`{
			"title": "firerice",
			"amount": 100,
			"note": "new dish",
			"tags": ["food"]
		}`)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/expenses", payload)

		//act
		r.ServeHTTP(w, req)
		got := responses.ExpenseResponse{}
		json.NewDecoder(w.Body).Decode(&got)

		//assert
		assert.Equal(t, http.StatusCreated, w.Code)
		if !assert.ObjectsAreEqual(want, got) {
			t.Errorf("not equal. want: %#v, got: %#v", want, got)
		}
	})

	t.Run("fail bad request because leave some json's field blank", func(t *testing.T) {
		//arrange
		expenseService := services.NewExpenseServiceMock()
		expenseHandler := handlers.NewExpenseHandler(expenseService)

		r := gin.Default()
		r.POST("/expenses", expenseHandler.CreateExpense)

		payload := strings.NewReader(`{
			"title": "firerice",
			"amount": 100,
			"note": "new dish"
		}`)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/expenses", payload)

		//act
		r.ServeHTTP(w, req)

		//assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("fail to create expense because internal server error", func(t *testing.T) {
		//arrange
		expenseService := services.NewExpenseServiceMock()
		expenseService.On("CreateExpense").Return(&responses.ExpenseResponse{}, helpers.NewInternalServerError())

		expenseHandler := handlers.NewExpenseHandler(expenseService)

		r := gin.Default()
		r.POST("/expenses", expenseHandler.CreateExpense)

		payload := strings.NewReader(`{
			"title": "firerice",
			"amount": 100,
			"note": "new dish",
			"tags": ["food"]
		}`)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/expenses", payload)

		//act
		r.ServeHTTP(w, req)

		//assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetExpenseByIDHandler(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		//arrange
		id := "1"
		want := responses.ExpenseResponse{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}

		expenseService := services.NewExpenseServiceMock()
		expenseService.On("GetExpenseByID", id).Return(&want, nil)

		expenseHandler := handlers.NewExpenseHandler(expenseService)

		r := gin.Default()
		r.GET("/expenses/:id", expenseHandler.GetExpenseByID)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/expenses/%s", id), nil)

		//act
		r.ServeHTTP(w, req)
		var got responses.ExpenseResponse
		json.NewDecoder(w.Body).Decode(&got)

		assert.Equal(t, http.StatusOK, w.Code)
		if !assert.ObjectsAreEqual(want, got) {
			t.Errorf("not equal. want: %#v, got: %#v", want, got)
		}

	})

	t.Run("fail case record not found", func(t *testing.T) {
		//arrange
		id := "1"

		expenseService := services.NewExpenseServiceMock()
		expenseService.On("GetExpenseByID", id).Return(&responses.ExpenseResponse{}, helpers.NewNotFoundError())

		expenseHandler := handlers.NewExpenseHandler(expenseService)

		r := gin.Default()
		r.GET("/expenses/:id", expenseHandler.GetExpenseByID)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/expenses/%s", id), nil)

		//act
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

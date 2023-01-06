//go:build integration
// +build integration

package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/wytquant/assessment/responses"
	"github.com/wytquant/assessment/src/expense/handlers"
	"github.com/wytquant/assessment/src/expense/repositories"
	"github.com/wytquant/assessment/src/expense/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var serverPort = 2565

func createAndSendReq(httpMethod string, url string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	return resp, err
}

func TestIntegrationTestServer(t *testing.T) {
	//setup server
	r := gin.Default()
	go func(r *gin.Engine) {
		db, err := gorm.Open(postgres.Open("postgres://root:root@db/go-integration-test-db?sslmode=disable"), &gorm.Config{})
		if err != nil {
			log.Fatalln("fail to connect the database")
		}

		repo := repositories.NewExpenseRepositoryDB(db)
		service := services.NewExpenseService(repo)
		handler := handlers.NewExpenseHandler(service)

		r.GET("/expenses/:id", handler.GetExpenseByID)
		r.GET("/expenses", handler.GetAllExpenses)
		r.POST("/expenses", handler.CreateExpense)
		r.PUT("/expenses/:id", handler.UpdateExpenseByID)

		r.Run(fmt.Sprintf(":%d", serverPort))

	}(r)

	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}
	t.Run("get all expense", func(t *testing.T) {
		//arrange and act
		resp, err := createAndSendReq(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses", serverPort), nil)
		assert.NoError(t, err)

		var got []responses.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		assert.NoError(t, err)
		resp.Body.Close()

		//assertion

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.NotEqual(t, 0, len(got))
		}
	})

	t.Run("get expense by id", func(t *testing.T) {
		//arrange
		id := 1

		//act
		resp, err := createAndSendReq(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses/%d", serverPort, id), nil)
		assert.NoError(t, err)

		var got responses.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		assert.NoError(t, err)
		resp.Body.Close()

		//assertion
		want := responses.ExpenseResponse{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			if !assert.ObjectsAreEqual(want, got) {
				t.Errorf("not equal. want: %#v but got: %#v", want, got)
			}
		}
	})

	t.Run("create expense", func(t *testing.T) {
		//arrange
		payload := strings.NewReader(`{
			"title": "strawberry smoothie",
			"amount": 79,
			"note": "night market promotion discount 10 bath",
			"tags": ["food", "beverage"]
		}`)

		//act
		resp, err := createAndSendReq(http.MethodPost, fmt.Sprintf("http://localhost:%d/expenses", serverPort), payload)
		assert.NoError(t, err)

		var got responses.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		assert.NoError(t, err)
		resp.Body.Close()

		//assertion
		want := responses.ExpenseResponse{
			ID:     2,
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food", "beverage"},
		}

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, resp.StatusCode)
			if !assert.ObjectsAreEqual(want, got) {
				t.Errorf("not equal. want: %#v but got: %#v", want, got)
			}
		}
	})

	t.Run("update expense by id", func(t *testing.T) {
		//arrange
		id := 1
		payload := strings.NewReader(`{
			"title": "strawberry smoothie",
			"amount": 100,
			"note": "night market promotion discount 10 bath",
			"tags": ["food"]
		}`)

		//act
		resp, err := createAndSendReq(http.MethodPut, fmt.Sprintf("http://localhost:%d/expenses/%d", serverPort, id), payload)
		assert.NoError(t, err)

		var got responses.ExpenseResponse
		err = json.NewDecoder(resp.Body).Decode(&got)
		assert.NoError(t, err)
		resp.Body.Close()

		//assertion
		want := responses.ExpenseResponse{
			ID:     1,
			Title:  "strawberry smoothie",
			Amount: 100,
			Note:   "night market promotion discount 10 bath",
			Tags:   pq.StringArray{"food"},
		}

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			if !assert.ObjectsAreEqual(want, got) {
				t.Errorf("not equal. want: %#v but got: %#v", want, got)
			}
		}
	})
}

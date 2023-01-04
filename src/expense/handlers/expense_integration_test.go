//go:build integration
// +build integration

package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
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

func TestIntegrationCreateExpense(t *testing.T) {
	//setup server
	r := gin.Default()
	go func(r *gin.Engine) {
		db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
		if err != nil {
			log.Fatalln("fail to connect the database")
		}

		repo := repositories.NewExpenseRepositoryDB(db)
		service := services.NewExpenseService(repo)
		handler := handlers.NewExpenseHandler(service)

		r.POST("/expenses", handler.CreateExpense)

		r.Run(fmt.Sprintf(":%d", 80))

	}(r)

	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", 80), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}

	//arrange
	payload := strings.NewReader(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%s/expenses", os.Getenv("PORT")), payload)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//act
	resp, err := client.Do(req)
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
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		if !assert.ObjectsAreEqual(want, got) {
			t.Errorf("not equal. want: %#v but got: %#v", want, got)
		}
	}
}

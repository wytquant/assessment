package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wytquant/assessment/config"
	"github.com/wytquant/assessment/src/expense/handlers"
	"github.com/wytquant/assessment/src/expense/repositories"
	"github.com/wytquant/assessment/src/expense/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authozired := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("USERNAME"): os.Getenv("PASSWORD"),
	}))

	{
		repo := repositories.NewExpenseRepositoryDB(config.DB)
		service := services.NewExpenseService(repo)
		expenseHandler := handlers.NewExpenseHandler(service)

		authozired.POST("/expenses", expenseHandler.CreateExpense)
		authozired.GET("/expenses/:id", expenseHandler.GetExpenseByID)
		authozired.PUT("/expenses/:id", expenseHandler.UpdateExpenseByID)
		authozired.GET("/expenses", expenseHandler.GetAllExpenses)
	}

	return r
}

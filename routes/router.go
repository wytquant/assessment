package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wytquant/assessment/config"
	"github.com/wytquant/assessment/src/expense/handlers"
	"github.com/wytquant/assessment/src/expense/repositories"
	"github.com/wytquant/assessment/src/expense/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	repo := repositories.NewExpenseRepositoryDB(config.DB)
	service := services.NewExpenseService(repo)
	expenseHandler := handlers.NewExpenseHandler(service)

	r.POST("/expenses", expenseHandler.CreateExpense)
	r.GET("/expenses/:id", expenseHandler.GetExpenseByID)
	r.PUT("/expenses/:id", expenseHandler.UpdateExpenseByID)

	return r
}

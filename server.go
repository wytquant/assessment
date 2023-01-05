package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/wytquant/assessment/config"
	"github.com/wytquant/assessment/models"
	"github.com/wytquant/assessment/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("fail to load .env file")
	}

	//connect to db
	err := config.InitPostgresDB()
	if err != nil {
		log.Fatalln("fail to connect the database")
	}
	defer config.CloseDB()

	config.DB.AutoMigrate(&models.Expense{})

	//setup routes
	r := routes.SetupRouter()

	//implement graceful shutdown
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

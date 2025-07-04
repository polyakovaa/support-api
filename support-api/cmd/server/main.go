package main

import (
	"askon/support-api/config"
	"askon/support-api/handlers"
	"askon/support-api/storage"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.Load("config.yaml")

	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	articleStorage := storage.NewArticleStorage(db)
	ticketStorage := storage.NewTicketStorage(db)
	handler := handlers.NewHandler(articleStorage, ticketStorage)
	webHandler := handlers.NewWebHandler()

	r := gin.Default()

	ticketGroup := r.Group("/api/tickets")
	{
		ticketGroup.GET("/states", handler.HandleTicketStates)
		ticketGroup.GET("/services", handler.HandleTicketServices)
	}

	articleGroup := r.Group("/api/articles")
	{
		articleGroup.GET("/types", handler.HandleArticleTypes)
		articleGroup.GET("/create-time", handler.HandleArticleTimes)
		articleGroup.GET("/senders", handler.HandleArticleSenders)
	}

	r.LoadHTMLGlob("support-api/web/templates/*")
	r.StaticFS("/static", gin.Dir("/app/support-api/web/static", true))
	r.GET("/dashboard", webHandler.ShowDashboard)

	log.Printf("Starting server on :%d", cfg.APIPort)
	if err := r.Run(fmt.Sprintf(":%d", cfg.APIPort)); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

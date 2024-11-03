package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	"speedy/internal/database"
	"speedy/internal/speedtest"
)

type Server struct {
	port int

	db *gorm.DB
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}
	config := database.Config{
		Schedule: "*/5 * * * *",
	}

	NewServer.db.FirstOrCreate(&database.Config{}, config)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	c := cron.New()

	c.AddFunc(config.Schedule, func() {
		speedtest.Speedtest("https://nbg1-speed.hetzner.com/100MB.bin")
	})
	c.Start()
	log.Printf("Setup Cronjobs. Current schedule: %v", c.Entries())

	return server
}

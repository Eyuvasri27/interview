package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"country-api/cache"
	"country-api/client"
	"country-api/handlers"
	"country-api/service"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize dependencies
	cache := cache.NewInMemoryCache()
	client := client.NewRestCountriesClient()
	service := service.NewCountryService(cache, client)
	handler := &handlers.CountryHandler{Service: service}

	// Create Echo instance
	e := echo.New()

	// Routes
	e.GET("/api/countries/search", handler.SearchCountry)

	// Start server
	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("Server stopped")
}

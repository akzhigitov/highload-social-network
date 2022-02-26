package main

import (
	"net/http"

	"github.com/akzhigitov/highload-social-network/src/routes"
	"github.com/akzhigitov/highload-social-network/src/store"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	store.Connect()
	defer store.Disconnect()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"*"},
	}))

	routes.Setup(e)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

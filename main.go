package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	getHostName := func() string {
		h, _ := os.Hostname()
		return h
	}

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("host: %v", getHostName()))
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: fmt.Sprintf("OK: %v", getHostName())})
	})

	e.POST("/upload", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, struct {
			Status string
		}{
			Status: fmt.Sprintf("Upload: %v", getHostName()),
		})
	})

	e.POST("/huge", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, struct {
			Status string
		}{
			Status: fmt.Sprintf("Huge Upload: %v", getHostName()),
		})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

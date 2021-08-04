package main

import (
	"github.com/blackNIKboard/go-examples/go-echo-server/shared"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.POST("/test", test)
	e.GET("/test1", test1)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler
func test(c echo.Context) error {
	var (
		//bodyBytes []byte
		request shared.Person
	)

	if err := c.Bind(&request); err != nil {
		return err
	}

	spew.Dump(request)

	return c.String(http.StatusOK, "Hello, TEST!")
}

// Handler
func test1(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, TY PIDOR!")
}

package main

import (
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Setup proxy
	url1, err := url.Parse("http://localhost:8081")
	if err != nil {
		e.Logger.Fatal(err)
	}

	rewrite := middleware.Rewrite(map[string]string{
		"/log/*": "/$1",
		"/":      "/kek",
	})

	g := e.Group("/log")
	g.Use(rewrite, middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{{URL: url1}})))
	g1 := e.Group("/")
	g1.Use(rewrite, middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{{URL: url1}})))
	e.Logger.Fatal(e.Start(":1323"))
}

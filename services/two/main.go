package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yoophi/go-monorepo/libs/hello"
)

func main() {
	e := echo.New()
	e.GET("/two/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, hello.Greet("Friend"))
	})
	_ = e.Start(":8080")
}

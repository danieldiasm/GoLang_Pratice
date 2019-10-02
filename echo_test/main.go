package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//Echo Instance
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Routes
	e.GET("/", hello)

	//Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

//Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hallo Welt!")
}

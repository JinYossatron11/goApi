package main

import (
	"projectapi/seven"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	//Load config

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/sevens", seven.CreateUser)
	e.GET("/sevens", seven.GetSeven)
	e.Logger.Fatal(e.Start(":1323"))
}

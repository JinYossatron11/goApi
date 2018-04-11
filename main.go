package main

import (
	"projectapi/seven"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	//Load config
	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/sevens", seven.CreateUser)
	e.GET("/sevens", seven.GetSeven)
	e.DELETE("/sevens/:id",seven.DeleteId)
	e.Logger.Fatal(e.Start(":1323"))
}

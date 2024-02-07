package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

// Handler
func getUsers(c echo.Context) error {
	u := &User{
	  Name:  "Jeff",
	  Email: "gholamzadeh.cuh@gmail.com",
	}
	return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", getUsers)
	e.Logger.Fatal(e.Start(":1323"))
}
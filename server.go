package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/article", createArticle)
	e.GET("/article/:id", getArticle)
	e.Logger.Fatal(e.Start(":1323"))
}

// create article interface
func createArticle(c echo.Context) error{
   content := c.Param("content")
   return c.String(http.StatusOK, content)
}

func getArticle(c echo.Context) error{
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}


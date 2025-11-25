package main

import (
	"net/http"
	// "os"
	// "time"
	// "strconv"
	// "math/rand"
	"sync/atomic"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID int `json: "id"`
	Title string `json: "title"`
	Completed bool `json: "completed"`
}
// Todo一覧用の箱
var todos []Todo
var nextID int64=0

func main() {
	e := echo.New()
	e.POST("/todos", func(c echo.Context) error {
		var todo Todo
		if err := c.Bind(&todo); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid json",
			})
		}
		// 自動採番
		newID := atomic.AddInt64(&nextID, 1)
		todo.ID = int(newID) 
		todos = append(todos, todo)

		return c.JSON(http.StatusOK, todo)
	})
	e.GET("/todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})
	e.Logger.Fatal(e.Start(":3080"))
}
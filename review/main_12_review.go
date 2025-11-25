package main

import (
	"fmt"
	"net/http"
	// "os"
	// "time"
	// "strconv"
	// "math/rand"
	"github.com/labstack/echo/v4"
	"sync/atomic"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Todo一覧用の箱
var todos []Todo
var nextID int64 = 0

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
		// review comment
		// ただ todos = append(todos, todo) は並列POSTが来ると本当は危ない（race）です。
		// 今はこのままで全然OK。
		// 本番だとここに Mutex か DB が要るんだな」くらいの理解で十分です。
		todos = append(todos, todo)

		return c.JSON(http.StatusOK, todo)
	})
	e.GET("/todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})
	e.Logger.Fatal(e.Start(":3080"))
}

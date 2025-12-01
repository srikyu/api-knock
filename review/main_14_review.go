package main

import (
	"net/http"
	// "os"
	// "time"
	"strconv"
	// "math/rand"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        int    `json: "id"`
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

// 更新用リクエストボディ
type TodoUpdateRequest struct {
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
		todos = append(todos, todo)

		return c.JSON(http.StatusOK, todo)
	})
	e.GET("/todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})
	e.GET("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		numID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "id must be an integer",
			})
		}
		for _, todo := range todos {
			if todo.ID == numID {
				return c.JSON(http.StatusOK, todo)
			}
		}
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "todo not found",
		})
	})
	e.PUT("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		numID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "id must be an integer",
			})
		}
		var req TodoUpdateRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid json",
			})
		}

		for _, todo := range todos {
			if todo.ID == numID {
				todo.Title = req.Title
				todo.Completed = req.Completed
				return c.JSON(http.StatusOK, todo)
			}
		}

		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "todo not found",
		})
	})
	e.Logger.Fatal(e.Start(":3080"))
}

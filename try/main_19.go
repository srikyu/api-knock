package main

import (
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

// Todo一覧用の箱
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// 更新用リクエストボディ
type TodoUpdateRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

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

		for i, todo := range todos {
			if todo.ID == numID {
				todos[i].Title = req.Title
				todos[i].Completed = req.Completed
				return c.JSON(http.StatusOK, todos[i])
			}
		}

		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "todo not found",
		})
	})
	e.PATCH("/todos/:id", func(c echo.Context) error {
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

		for i, todo := range todos {
			if todo.ID == numID {
				todos[i].Completed = req.Completed
				return c.JSON(http.StatusOK, todos[i])
			}
		}

		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "todo not found",
		})
	})
	e.DELETE("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		numID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "id must be an integer",
			})
		}

		for i, todo := range todos {
			if todo.ID == numID {
				todos = append(todos[:i], todos[i+1:]...)
				return c.JSON(http.StatusOK, map[string]string{
					"deleted": strconv.Itoa(numID),
				})
			}
		}
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "todo not found",
		})
	})

	e.DELETE("/todos", func(c echo.Context) error {
		force := c.QueryParam("force")
		if force == "true" {
			todos = todos[:0]
			return c.JSON(http.StatusOK, map[string]string{
				"deleted": "all",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"stats": "invalied parameter",
		})
	})

	e.GET("/todos", func(c echo.Context) error {
		params := c.QueryParam("completed")
		if params == "true" {
			var completedTodos []Todo
			for _, todo := range todos {
				if todo.Completed {
					completedTodos = append(completedTodos, todo)
				}
			}
			return c.JSON(http.StatusOK, completedTodos)
		}
		return c.JSON(http.StatusOK, todos)
	})

	e.POST("/todos/:id/complete", func(c echo.Context) error {
		id := c.Param("id")
		numID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "id must be an integer",
			})
		}

		for i, todo := range todos {
			if todo.ID == numID {
				todos[i].Completed = true
			}
		}
		return c.JSON(http.StatusOK, map[string]string{
			"id":        strconv.Itoa(numID),
			"completed": "true",
		})
	})
	e.Logger.Fatal(e.Start(":3080"))
}

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


var nextID int64=0

func main() {
	e := echo.New()

	//before review

	// type Todo struct {
	// 	ID int `json: id`
	// 	Title string `json: title`
	// 	Completed bool `json: completed`
	// }
	// // 11. POST /todos
	// e.POST("/todos", func(c echo.Context) error {
	// 	var todo Todo
	// 	if err := c.Bind(&todo); err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"id": strconv.Itoa(todo.ID),
	// 		"title": todo.Title,
	// 		"completed": strconv.FormatBool(todo.Completed),
	// 	})
	// })

	//after review
	type Todo struct {
		ID int `json: "id"`
		Title string `json: "title"`
		Completed bool `json: "completed"`
	}

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

		return c.JSON(http.StatusOK, todo)
	})
	e.Logger.Fatal(e.Start(":3080"))
}
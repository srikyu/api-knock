package main

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	e := echo.New()

	// 1.ping
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	// 2. hello
	e.GET("/hello", func(c echo.Context) error {
		name := c.QueryParam("name")
		if name == "" {
			name = "world"
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello " + name,
		})
	})

	type EchoBody struct {
		Message string `json:"message"`
	}

	// 3. POST /echo -> bodyをそのまま返す
	e.POST("/echo", func(c echo.Context) error {
		var req EchoBody
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid json",
			})
		}
		// そのまま返す
		return c.JSON(http.StatusOK, req)
	})

	// 4. TIME /time
	e.GET("/time", func(c echo.Context) error {
		now := time.Now().Format(time.RFC3339)
		return c.JSON(http.StatusOK, map[string]string{
			"time": now,
		})
	})

	// 5. GET /env
	e.GET("/env", func(c echo.Context) error {
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "unknown"
		}
		return c.JSON(http.StatusOK, map[string]string{
			"app_env": env,
		})
	})

	// 6. GET /headers
	e.GET("/headers", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.Request().Header)
	})

	// 7. GET /status/:code
	e.GET("/status/:code", func(c echo.Context) error {
		codeStr := c.Param("code")
		code, err := strconv.Atoi(codeStr)
		if err != nil || code < 100 || code > 599 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid status code",
			})
		}
		return c.JSON(code, map[string]interface{}{
			"status": code,
		})
	})

	// 8. GET /random/
	e.GET("/random/int", func(c echo.Context) error {
		min := c.QueryParam("min")
		max := c.QueryParam("max")

		if min == "" {
			min = "0"
		}
		if max == "" {
			max = "100"
		}
		minInt, err := strconv.Atoi(min)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid min value",
			})
		}
		maxInt, err := strconv.Atoi(max)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid max value",
			})
		}
		n := rand.Intn(maxInt-minInt+1) + minInt
		return c.JSON(http.StatusOK, map[string]int{
			"random": n,
		})
	})

	// 9. GET /fib/:n
	e.GET("/fib/:n", func(c echo.Context) error {
		nStr := c.Param("n")
		n, err := strconv.Atoi(nStr)
		// nに関するエラーハンドリング
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid n value",
			})
		}

		// fib関数を定義
		var fib func(int) int
		fib = func(n int) int {
			if n < 2 {
				return n
			}
			return fib(n-1) + fib(n-2)
		}

		result := fib(n)
		if result >= 400 {
			return c.JSON(http.StatusBadRequest, map[string]int{
				"fib": 400,
			})
		}

		return c.JSON(http.StatusOK, map[string]int{
			"fib": result,
		})
	})

	// 10. GET /health
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})
	e.Logger.Fatal(e.Start(":3080"))
}

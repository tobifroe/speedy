package server

import (
	"fmt"
	"net/http"
	"speedy/internal/database"
	"speedy/internal/speedtest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)

	e.GET("/speed", s.speedHandler)

	e.GET("/recent", s.mostRecentresult)
	e.GET("/all", s.allResults)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) speedHandler(c echo.Context) error {
	elapsed := speedtest.Speedtest("https://nbg1-speed.hetzner.com/100MB.bin")

	resp := map[string]string{
		"message": "Hello ",
		"elapsed": elapsed.String(),
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) mostRecentresult(c echo.Context) error {
	var testresult database.TestResult
	s.db.First(&testresult)

	duration := fmt.Sprintf("%f", testresult.Duration)

	resp := map[string]string{
		"message":  "DB Read",
		"target":   testresult.Target,
		"duration": duration,
	}

	return c.JSON(http.StatusOK, resp)
}

type allResults struct {
	Results []database.TestResult
}

func (s *Server) allResults(c echo.Context) error {

	var testresults []database.TestResult
	s.db.Find(&testresults)

	resp := allResults{
		Results: testresults,
	}
	return c.JSON(http.StatusOK, resp)
}

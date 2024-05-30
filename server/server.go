package server

import (
	"capstone-mikti/configs"
	"capstone-mikti/utils/database"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo
	c *configs.ProgrammingConfig
}

func (s *Server) RunServer() {
	s.e.Pre(middleware.RemoveTrailingSlash())

	s.e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	s.e.Use(middleware.Recover())

	s.e.Logger.Debug()
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%d", s.c.Server)).Error())
}

func (s *Server) MigrateDB() {
	db := database.InitDB(s.c)
	database.Migrate(db)
}

func InitServer(e *echo.Echo, c *configs.ProgrammingConfig) *Server {
	return &Server{
		e: e,
		c: c,
	}
}

package server

import (
	"capstone-mikti/configs"
	"capstone-mikti/utils/database"
	"capstone-mikti/utils/database/seeds"
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

func (s *Server) SeederDB() {
	db := database.InitDB(s.c)
	// HANYA 1x SAJA
	for _, seed := range seeds.All() {
		if err := seed.Run(db); err != nil {
			fmt.Printf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}

func InitServer(e *echo.Echo, c *configs.ProgrammingConfig) *Server {
	return &Server{
		e: e,
		c: c,
	}
}

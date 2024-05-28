package routes

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRoute(c *configs.ProgrammingConfig, uh users.UserHandlerInterface) *echo.Echo {
	e := echo.New()

	group := e.Group("/api/v1")

	// Route Authentication
	group.POST("/register", uh.Register())
	group.POST("/login", uh.Login())
	group.POST("/forget-password", uh.ForgetPasswordWeb())
	group.POST("/reset-password", uh.ResetPassword())
	group.POST("/refresh-token", uh.RefreshToken(), echojwt.JWT([]byte(c.Secret)))

	// Route Profile
	group.GET("/profile", uh.Profile(), echojwt.JWT([]byte(c.Secret)))
	group.POST("/profile/update", uh.UpdateProfile(), echojwt.JWT([]byte(c.Secret)))

	return e
}

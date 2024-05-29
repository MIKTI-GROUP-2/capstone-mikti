package routes

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/users"
	"capstone-mikti/helper/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute(c *configs.ProgrammingConfig, uh users.UserHandlerInterface) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	config := configs.InitConfig()

	//Akses khusus sesuai role
	RoleMiddleware := middlewares.NewMiddleware(config)

	//Akses khusus harus login dlu
	JwtAuth := echojwt.JWT([]byte(c.Secret))

	group := e.Group("/api/v1")

	// Route Authentication
	group.POST("/register", uh.Register())
	group.POST("/login", uh.Login())
	group.POST("/forget-password", uh.ForgetPasswordWeb())
	group.POST("/reset-password", uh.ResetPassword())
	group.POST("/refresh-token", uh.RefreshToken(), JwtAuth)

	// Route Profile
	group.GET("/profile", uh.Profile(), JwtAuth)
	group.POST("/profile/update", uh.UpdateProfile(), JwtAuth)

	// Route Group event
	groupEvent := group.Group("/event")
	groupEvent.Use(RoleMiddleware.RoleMiddleware("Admin"))
	groupEvent.GET("", uh.Profile())

	return e
}

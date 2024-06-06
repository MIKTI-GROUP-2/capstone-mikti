package routes

import (
	"capstone-mikti/configs"

	"capstone-mikti/features/categories"
	events "capstone-mikti/features/events"
	"capstone-mikti/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRoute(c *configs.ProgrammingConfig, uh users.UserHandlerInterface, eh events.EventHandlerInterface, ch categories.CategoryHandlerInterface) *echo.Echo {
	e := echo.New()

	//Akses khusus harus login dlu
	JwtAuth := echojwt.JWT([]byte(c.Secret))

	group := e.Group("/api/v1")

	// Route Authentication
	group.POST("/register", uh.Register())
	group.POST("/login", uh.Login())
	group.POST("/forget-password", uh.ForgetPasswordWeb())
	group.POST("/reset-password", uh.ResetPassword())
	group.POST("/refresh-token", uh.RefreshToken(), JwtAuth)
	group.POST("/refresh-token", uh.RefreshToken(), JwtAuth)

	// Route Profile
	group.GET("/profile", uh.Profile(), JwtAuth)
	group.POST("/profile/update", uh.UpdateProfile(), JwtAuth)

	//Route Event Category
	group.GET("/categories", ch.GetCategories(), JwtAuth)
	group.GET("/category/:id", ch.GetCategory(), JwtAuth)
	group.POST("/category", ch.CreateCategory(), JwtAuth)
	group.PUT("/category/:id", ch.UpdateCategory(), JwtAuth)

	// Route Group event
	groupEvent := group.Group("/event")
	groupEvent.GET("", eh.GetAll())
	groupEvent.POST("", eh.CreateEvent(), JwtAuth)
	groupEvent.GET("/:id", eh.GetDetail())
	groupEvent.PUT("/:id", eh.UpdateEvent(), JwtAuth)
	groupEvent.DELETE("/:id", eh.DeleteEvent(), JwtAuth)

	return e
}

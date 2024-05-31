package routes

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/users"
	"capstone-mikti/features/wishlists"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRoute(c *configs.ProgrammingConfig, uh users.UserHandlerInterface, wh wishlists.WishlistHandlerInterface) *echo.Echo {
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

	// Route Profile
	group.GET("/profile", uh.Profile(), JwtAuth)
	group.POST("/profile/update", uh.UpdateProfile(), JwtAuth)

	// Route Group event
	groupEvent := group.Group("/event")
	groupEvent.GET("", uh.Profile(), JwtAuth)

	// Route Wishlist
	groupWishlist := group.Group("/wishlist")
	groupWishlist.GET("", wh.GetAll())
	groupWishlist.GET("/:user_id", wh.GetByUserID())
	groupWishlist.POST("/create", wh.Create())
	groupWishlist.DELETE("/:id", wh.Delete())

	return e
}

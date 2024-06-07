package routes

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/tickets"
	"capstone-mikti/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRoute(c *configs.ProgrammingConfig, uh users.UserHandlerInterface, th tickets.TicketHandlerInterface) *echo.Echo {
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

	// Route Ticket
	groupTicket := group.Group("/ticket")
	groupTicket.POST("", th.Create(), JwtAuth)
	groupTicket.GET("", th.GetAll(), JwtAuth)
	groupTicket.GET("/:id", th.GetByID(), JwtAuth)
	groupTicket.PUT("/:id", th.Update(), JwtAuth)

	return e
}

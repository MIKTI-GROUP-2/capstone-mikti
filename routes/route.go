package routes

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/payments"
	"capstone-mikti/features/tickets"

	"capstone-mikti/features/bookings"
	"capstone-mikti/features/categories"
	events "capstone-mikti/features/events"
	"capstone-mikti/features/users"
	"capstone-mikti/features/vouchers"
	"capstone-mikti/features/wishlists"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRoute(c *configs.ProgrammingConfig, uh users.UserHandlerInterface, eh events.EventHandlerInterface, ch categories.CategoryHandlerInterface, wh wishlists.WishlistHandlerInterface, th tickets.TicketHandlerInterface, vh vouchers.VoucherHandlerInterface, bh bookings.BookingHandlerInterface, ph payments.PaymentHandlerInterface) *echo.Echo {
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
	group.GET("/category", ch.GetCategories(), JwtAuth)
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
	groupEvent.GET("/:event_id/ticket", th.GetByEventID())

	// Route Group Voucher
	groupVoucher := group.Group("/voucher")
	groupVoucher.GET("", vh.GetVouchers())
	groupVoucher.GET("/:id", vh.GetVoucher())
	groupVoucher.POST("", vh.CreateVoucher(), JwtAuth)
	groupVoucher.PUT("/:id", vh.UpdateVoucher(), JwtAuth)
	groupVoucher.GET("/:id/activate", vh.ActivateVoucher(), JwtAuth)
	groupVoucher.GET("/:id/deactivate", vh.DeactivateVoucher(), JwtAuth)

	// Route Wishlist
	groupWishlist := group.Group("/wishlist")
	groupWishlist.POST("", wh.Create(), JwtAuth)
	groupWishlist.GET("", wh.GetAll(), JwtAuth)
	groupWishlist.GET("/:event_id", wh.GetByEventID(), JwtAuth)
	groupWishlist.DELETE("/:event_id", wh.Delete(), JwtAuth)

	// Route Ticket
	groupTicket := group.Group("/ticket")
	groupTicket.POST("", th.Create(), JwtAuth)
	groupTicket.GET("", th.GetAll())
	groupTicket.GET("/:id", th.GetByID())
	groupTicket.PUT("/:id", th.Update(), JwtAuth)
	groupTicket.DELETE("/:id", th.Delete(), JwtAuth)

	//Booking
	groupBooking := group.Group("/booking")
	groupBooking.POST("", bh.CreateBooking(), JwtAuth)
	groupBooking.GET("", bh.GetAll(), JwtAuth)
	groupBooking.GET("/:id", bh.GetDetail(), JwtAuth)
	groupBooking.DELETE("/:id", bh.DeleteBooking(), JwtAuth)

	groupPayment := group.Group("/payment")
	groupPayment.GET("", ph.GetAll(), JwtAuth)
	groupPayment.POST("", ph.CreatePayment(), JwtAuth)
	groupPayment.POST("/notif", ph.NotifPayment())

	return e
}

//go:build wireinject
// +build wireinject

package main

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/bookings"
	"capstone-mikti/features/categories"
	"capstone-mikti/features/events"
	"capstone-mikti/features/users"
	"capstone-mikti/features/vouchers"
	"capstone-mikti/features/wishlists"
	"capstone-mikti/helper/email"
	"capstone-mikti/helper/enkrip"
	"capstone-mikti/helper/jwt"
	"capstone-mikti/routes"
	"capstone-mikti/utils/cloudinary"
	"capstone-mikti/utils/database"

	userData "capstone-mikti/features/users/data"
	userHandler "capstone-mikti/features/users/handler"
	userService "capstone-mikti/features/users/service"

	wishlistData "capstone-mikti/features/wishlists/data"
	wishlistHandler "capstone-mikti/features/wishlists/handler"
	wishlistService "capstone-mikti/features/wishlists/service"

	eventData "capstone-mikti/features/events/data"
	eventHandler "capstone-mikti/features/events/handler"
	eventService "capstone-mikti/features/events/service"

	categoryData "capstone-mikti/features/categories/data"
	categoryHandler "capstone-mikti/features/categories/handler"
	categoryService "capstone-mikti/features/categories/service"

	voucherData "capstone-mikti/features/vouchers/data"
	voucherHandler "capstone-mikti/features/vouchers/handler"
	voucherService "capstone-mikti/features/vouchers/service"

	bookingData "capstone-mikti/features/bookings/data"
	bookingHandler "capstone-mikti/features/bookings/handler"
	bookingService "capstone-mikti/features/bookings/service"

	"capstone-mikti/server"

	"github.com/google/wire"
)

var userSet = wire.NewSet(
	userData.New,
	wire.Bind(new(users.UserDataInterface), new(*userData.UserData)),

	userService.New,
	wire.Bind(new(users.UserServiceInterface), new(*userService.UserService)),

	userHandler.NewHandler,
	wire.Bind(new(users.UserHandlerInterface), new(*userHandler.UserHandler)),
)
var categorySet = wire.NewSet(
	categoryData.New,
	wire.Bind(new(categories.CategoryDataInterface), new(*categoryData.CategoryData)),

	categoryService.New,
	wire.Bind(new(categories.CategoryServiceInterface), new(*categoryService.CategoryService)),

	categoryHandler.NewHandler,
	wire.Bind(new(categories.CategoryHandlerInterface), new(*categoryHandler.CategoryHandler)),
)

var eventSet = wire.NewSet(
	eventData.New,
	wire.Bind(new(events.EventDataInterface), new(*eventData.EventData)),

	eventService.New,
	wire.Bind(new(events.EventServiceInterface), new(*eventService.EventService)),

	eventHandler.NewHandler,
	wire.Bind(new(events.EventHandlerInterface), new(*eventHandler.EventHandler)),
)

var voucherSet = wire.NewSet(
	voucherData.New,
	wire.Bind(new(vouchers.VoucherDataInterface), new(*voucherData.VoucherData)),

	voucherService.New,
	wire.Bind(new(vouchers.VoucherServiceInterface), new(*voucherService.VoucherService)),

	voucherHandler.NewHandler,
	wire.Bind(new(vouchers.VoucherHandlerInterface), new(*voucherHandler.VoucherHandler)),
)

var wishlistSet = wire.NewSet(
	wishlistData.New,
	wire.Bind(new(wishlists.WishlistDataInterface), new(*wishlistData.WishlistData)),

	wishlistService.New,
	wire.Bind(new(wishlists.WishlistServiceInterface), new(*wishlistService.WishlistService)),

	wishlistHandler.NewHandler,
	wire.Bind(new(wishlists.WishlistHandlerInterface), new(*wishlistHandler.WishlistHandler)),
)

var bookingSet = wire.NewSet(
	bookingData.New,
	wire.Bind(new(bookings.BookingDataInterface), new(*bookingData.BookingData)),

	bookingService.New,
	wire.Bind(new(bookings.BookingServiceInterface), new(*bookingService.BookingService)),

	bookingHandler.NewHandler,
	wire.Bind(new(bookings.BookingHandlerInterface), new(*bookingHandler.BookingHandler)),
)

func InitializedServer() *server.Server {
	wire.Build(
		configs.InitConfig,
		database.InitDB,
		enkrip.New,
		email.New,
		jwt.NewJWT,
		cloudinary.InitCloud,
		// JANGAN DIRUBAH

		userSet,
		eventSet,
		categorySet,
		voucherSet,
		wishlistSet,
		bookingSet,

		// JANGAN DIRUBAH
		routes.NewRoute,
		server.InitServer,
	)

	return nil
}

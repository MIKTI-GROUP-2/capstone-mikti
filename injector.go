//go:build wireinject
// +build wireinject

package main

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/users"
	"capstone-mikti/features/wishlists"
	"capstone-mikti/helper/email"
	"capstone-mikti/helper/enkrip"
	"capstone-mikti/helper/jwt"
	"capstone-mikti/routes"
	"capstone-mikti/utils/database"

	userData "capstone-mikti/features/users/data"
	userHandler "capstone-mikti/features/users/handler"
	userService "capstone-mikti/features/users/service"

	wishlistData "capstone-mikti/features/wishlists/data"
	wishlistHandler "capstone-mikti/features/wishlists/handler"
	wishlistService "capstone-mikti/features/wishlists/service"

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

var wishlistSet = wire.NewSet(
	wishlistData.New,
	wire.Bind(new(wishlists.WishlistDataInterface), new(*wishlistData.WishlistData)),

	wishlistService.New,
	wire.Bind(new(wishlists.WishlistServiceInterface), new(*wishlistService.WishlistService)),

	wishlistHandler.NewHandler,
	wire.Bind(new(wishlists.WishlistHandlerInterface), new(*wishlistHandler.WishlistHandler)),
)

func InitializedServer() *server.Server {
	wire.Build(
		configs.InitConfig,
		database.InitDB,
		enkrip.New,
		email.New,
		jwt.NewJWT,
		// JANGAN DIRUBAH

		userSet,
		wishlistSet,

		// JANGAN DIRUBAH
		routes.NewRoute,
		server.InitServer,
	)

	return nil
}

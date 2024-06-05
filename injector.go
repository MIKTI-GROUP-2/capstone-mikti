//go:build wireinject
// +build wireinject

package main

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/categories"
	"capstone-mikti/features/events"
	"capstone-mikti/features/users"
	"capstone-mikti/helper/email"
	"capstone-mikti/helper/enkrip"
	"capstone-mikti/helper/jwt"
	"capstone-mikti/routes"
	"capstone-mikti/utils/cloudinary"
	"capstone-mikti/utils/database"

	userData "capstone-mikti/features/users/data"
	userHandler "capstone-mikti/features/users/handler"
	userService "capstone-mikti/features/users/service"

	eventData "capstone-mikti/features/events/data"
	eventHandler "capstone-mikti/features/events/handler"
	eventService "capstone-mikti/features/events/service"

	categoryData "capstone-mikti/features/categories/data"
	categoryHandler "capstone-mikti/features/categories/handler"
	categoryService "capstone-mikti/features/categories/service"

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

		// JANGAN DIRUBAH
		routes.NewRoute,
		server.InitServer,
	)

	return nil
}

//go:build wireinject
// +build wireinject

package main

import (
	"capstone-mikti/configs"
	"capstone-mikti/features/tickets"
	"capstone-mikti/features/users"
	"capstone-mikti/helper/email"
	"capstone-mikti/helper/enkrip"
	"capstone-mikti/helper/jwt"
	"capstone-mikti/routes"
	"capstone-mikti/utils/database"

	userData "capstone-mikti/features/users/data"
	userHandler "capstone-mikti/features/users/handler"
	userService "capstone-mikti/features/users/service"

	ticketData "capstone-mikti/features/tickets/data"
	ticketHandler "capstone-mikti/features/tickets/handler"
	ticketService "capstone-mikti/features/tickets/service"

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

var ticketSet = wire.NewSet(
	ticketData.New,
	wire.Bind(new(tickets.TicketDataInterface), new(*ticketData.TicketData)),

	ticketService.New,
	wire.Bind(new(tickets.TicketServiceInterface), new(*ticketService.TicketService)),

	ticketHandler.NewHandler,
	wire.Bind(new(tickets.TicketHandlerInterface), new(*ticketHandler.TicketHandler)),
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
		ticketSet,

		// JANGAN DIRUBAH
		routes.NewRoute,
		server.InitServer,
	)

	return nil
}

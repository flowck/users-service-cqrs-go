package main

import (
	"fmt"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/app/command"
	"users-service-cqrs/internal/app/query"
)

func main() {
	application := app.App{
		Commands: &app.Commands{
			BlockUser:   command.NewBlockUserHandler(nil),
			UnBlockUser: command.NewUnblockUserHandler(nil),
		},
		Queries: &app.Queries{
			AllBlockedUser: query.NewAllBlockedUsersHandler(nil),
		},
	}

	fmt.Println(application)
}

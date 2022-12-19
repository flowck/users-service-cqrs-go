package main

import (
	"context"
	"fmt"
	"users-service-cqrs/internal/adapters"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/app/command"
	"users-service-cqrs/internal/app/query"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//
	// Data Repositories
	//
	inMemReadUserRepo := adapters.NewInMemoryReadUserRepository()
	inMemWriteUserRepo := adapters.NewInMemoryWriteUserRepository()

	//
	// Assemble the application
	//
	application := app.App{
		Commands: &app.Commands{
			BlockUser:   command.NewBlockUserHandler(inMemWriteUserRepo),
			UnBlockUser: command.NewUnblockUserHandler(inMemWriteUserRepo),
		},
		Queries: &app.Queries{
			AllBlockedUser: query.NewAllBlockedUsersHandler(inMemReadUserRepo),
		},
	}

	fmt.Println(application.Queries.AllBlockedUser.Handle(ctx, query.AllBlockedUsers{}))
}

package app

import (
	"database/sql"
	"users-service-cqrs/internal/adapters"
	"users-service-cqrs/internal/app/command"
	"users-service-cqrs/internal/app/query"
)

type App struct {
	Queries  *Queries
	Commands *Commands
}

type Queries struct {
	AllUsers query.AllUsersHandler
	OneUser  query.OneUserHandler
}

type Commands struct {
	BlockUser   command.BlockUserHandler
	UnBlockUser command.UnBlockUserHandler
}

func New(db *sql.DB) *App {
	readUserRepo := adapters.NewPsqlReadUserRepo(db)
	writeUserRepo := adapters.NewPsqlWriteUserRepo(db)

	return &App{
		Commands: &Commands{
			BlockUser:   command.NewBlockUserHandler(writeUserRepo),
			UnBlockUser: command.NewUnblockUserHandler(writeUserRepo),
		},
		Queries: &Queries{
			AllUsers: query.NewAllBlockedUsersHandler(readUserRepo),
			OneUser:  query.NewOneUserHandler(readUserRepo),
		},
	}
}

package app

import (
	"users-service-cqrs/internal/app/command"
	"users-service-cqrs/internal/app/query"
)

type App struct {
	Queries  *Queries
	Commands *Commands
}

type Queries struct {
	AllUsers query.AllUsersHandler
	OneUser  query.AllUsersHandler
}

type Commands struct {
	BlockUser   command.BlockUserHandler
	UnBlockUser command.UnBlockUserHandler
}

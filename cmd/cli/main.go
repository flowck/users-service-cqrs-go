package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/app/query"
	"users-service-cqrs/internal/common/config"
	"users-service-cqrs/internal/common/psql"

	_ "github.com/lib/pq"
)

const (
	BlockUser int = iota
	UnBlockUser
	ShowUsers
)

func printInstructions() {
	fmt.Println("Please type of of the available commands listed below:")
	fmt.Printf("BlockUser:   %d\nUnBlockUser: %d\nShowUsers:   %d\n", BlockUser, UnBlockUser, ShowUsers)
}

func main() {
	cfg := config.New()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db := psql.Connect(cfg.PsqlUri)

	printInstructions()

	application := app.New(db)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("%s isn't a valid command", scanner.Text())
		}

		switch cmd {
		case BlockUser:
			fmt.Println("Please provide the user's id")
		case ShowUsers:
			fmt.Println("Please type the status: blocked or unblocked")
			statusScan := bufio.NewScanner(os.Stdin)
			for statusScan.Scan() {
				status := statusScan.Text()

				if status != "blocked" && status != "unblocked" {
					log.Fatalf("The status provided is invalid: %s", status)
				}

				userList, err := application.Queries.AllUsers.Handle(ctx, query.AllUsers{Status: status})
				if err != nil {
					log.Fatalf("unable to show users due to: %v", err)
				}

				for _, u := range userList {
					fmt.Println(u)
				}

				fmt.Print("\n")
				break
			}

			printInstructions()
		default:
			fmt.Printf("%d isn't a valid command\n\n", cmd)
			printInstructions()
		}
	}

	fmt.Println("End")
}

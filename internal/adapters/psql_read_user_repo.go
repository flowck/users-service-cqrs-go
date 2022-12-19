package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"users-service-cqrs/internal/adapters/models"
	"users-service-cqrs/internal/domain/user"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type psqlReadUserRepo struct {
	db *sql.DB
}

func (p psqlReadUserRepo) FindAll(ctx context.Context) ([]*user.User, error) {
	rows, err := models.Users(qm.Limit(100)).All(ctx, p.db)

	fmt.Println(rows[0].Email)

	if err != nil {
		return nil, err
	}

	return mapToUserList(rows)
}

func (p psqlReadUserRepo) Find(ctx context.Context, id *user.ID) (*user.User, error) {
	row, err := models.FindUser(ctx, p.db, id.String())
	if err != nil {
		return nil, err
	}

	return mapToUser(row)
}

func NewPsqlReadUserRepo(db *sql.DB) *psqlReadUserRepo {
	if db == nil {
		panic("db is nil")
	}

	return &psqlReadUserRepo{db: db}
}

func mapToUser(row *models.User) (*user.User, error) {
	id, err := user.NewIDFromString(row.ID)
	if err != nil {
		return nil, err
	}

	email, err := user.NewEmail(row.Email)
	if err != nil {
		return nil, err
	}

	u, err := user.New(id, row.FirstName.String, row.LastName.String, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func mapToUserList(rows []*models.User) ([]*user.User, error) {
	userList := make([]*user.User, len(rows))

	for idx, row := range rows {
		u, err := mapToUser(row)
		if err != nil {
			return nil, err
		}

		userList[idx] = u
	}

	return userList, nil
}

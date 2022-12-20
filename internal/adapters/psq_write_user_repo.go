package adapters

import (
	"context"
	"database/sql"
	"errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"users-service-cqrs/internal/adapters/models"
	"users-service-cqrs/internal/domain/user"
)

type psqlWriteUserRepo struct {
	db *sql.DB
}

var _ user.WriteRepository = (*psqlWriteUserRepo)(nil)

func (p psqlWriteUserRepo) Update(ctx context.Context, id *user.ID, updateFn func(u *user.User) *user.User) error {
	row, err := models.FindUser(ctx, p.db, id.String())
	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if err != nil {
		return err
	}

	u, err := mapToDomainUser(row)
	if err != nil {
		return err
	}

	updatedUser := updateFn(u)

	row.Email = updatedUser.Email().String()
	row.LastName = null.StringFrom(updatedUser.LastName())
	row.FirstName = null.StringFrom(updatedUser.FirstName())

	if updatedUser.IsBlocked() {
		row.Status = "blocked"
	} else {
		row.Status = "unblocked"
	}

	if _, err = row.Update(ctx, p.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func NewPsqlWriteUserRepo(db *sql.DB) *psqlWriteUserRepo {
	if db == nil {
		panic("db is nil")
	}

	return &psqlWriteUserRepo{db: db}
}

func mapToDomainUser(row *models.User) (*user.User, error) {
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

	if row.Status == "blocked" {
		u.Block()
	}

	if row.Status == "unblocked" {
		u.UnBlock()
	}

	return u, nil
}

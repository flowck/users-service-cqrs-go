package adapters

import (
	"context"
	"database/sql"
	"errors"
	"users-service-cqrs/internal/adapters/models"
	"users-service-cqrs/internal/app/query"
	"users-service-cqrs/internal/domain/user"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type psqlReadUserRepo struct {
	db *sql.DB
}

var _ query.ReadRepository = (*psqlReadUserRepo)(nil)

func (p psqlReadUserRepo) FindAll(ctx context.Context, q query.AllUsers) ([]*query.User, error) {
	rows, err := models.Users(qm.Limit(100), models.UserWhere.Status.EQ(q.Status)).All(ctx, p.db)

	if err != nil {
		return nil, err
	}

	return mapToUserList(rows)
}

func (p psqlReadUserRepo) Find(ctx context.Context, id *user.ID) (*query.User, error) {
	row, err := models.FindUser(ctx, p.db, id.String())

	if errors.Is(err, sql.ErrNoRows) {
		return nil, user.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return mapToUser(row), nil
}

func NewPsqlReadUserRepo(db *sql.DB) *psqlReadUserRepo {
	if db == nil {
		panic("db is nil")
	}

	return &psqlReadUserRepo{db: db}
}

func mapToUser(row *models.User) *query.User {
	return &query.User{
		Id:        row.ID,
		FirstName: row.FirstName.String,
		LastName:  row.LastName.String,
		Email:     row.Email,
		Status:    row.Status,
	}
}

func mapToUserList(rows []*models.User) ([]*query.User, error) {
	userList := make([]*query.User, len(rows))

	for idx, row := range rows {
		userList[idx] = mapToUser(row)
	}

	return userList, nil
}

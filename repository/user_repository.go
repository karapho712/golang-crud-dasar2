package repository

import (
	"context"
	"crud-dasar-go-2/model/entity"
	"database/sql"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error)
}

package impl

import (
	"context"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/entity"
	"crud-dasar-go-2/repository"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

// FindByEmail implements repository.UserRepository.
func (userRepostiory *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.User, error) {
	SQL := "SELECT id, name, email, password FROM user WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)

	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

// Save implements repository.UserRepository.
func (userRepostiory *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := "INSERT INTO user(name, password, email) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Password, user.Email)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}

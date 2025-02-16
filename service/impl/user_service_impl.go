package impl

import (
	"context"
	"crud-dasar-go-2/exception"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/entity"
	"crud-dasar-go-2/model/web/user"
	"crud-dasar-go-2/repository"
	"crud-dasar-go-2/service"
	"database/sql"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(
	UserRepository repository.UserRepository,
	DB *sql.DB,
	Validate *validator.Validate,
) service.UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		DB:             DB,
		Validate:       Validate,
	}
}

// Create implements service.UserService.
func (userService *UserServiceImpl) Create(ctx context.Context, request user.UserCreateRequest) user.UserResponse {
	err := userService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		helper.PanicIfError(err)
	}
	request.Password = string(bytes)

	user := entity.User{
		Name:     request.Name,
		Password: request.Password,
		Email:    request.Email,
	}
	user = userService.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

// FindByEmailAndPassword implements service.UserService.
func (userService *UserServiceImpl) FindByEmailAndPassword(ctx context.Context, request user.UserLoginRequest) user.UserResponse {
	err := userService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser, err := userService.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		// panic(exception.NewNotFoundError(err.Error()))
		panic(exception.NewUnauthorizedError(err.Error()))

	}

	comparedErr := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(request.Password))
	if comparedErr != nil {
		panic(exception.NewUnauthorizedError(comparedErr.Error()))
	}

	return helper.ToUserResponse(getUser)
}

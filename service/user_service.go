package service

import (
	"context"
	"crud-dasar-go-2/model/web/user"
)

type UserService interface {
	Create(ctx context.Context, request user.UserCreateRequest) user.UserResponse
	FindByEmailAndPassword(ctx context.Context, request user.UserLoginRequest) user.UserResponse
}

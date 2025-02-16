package user

type UserLoginRequest struct {
	Email    string `validate:"required,max=100,min=3" json:"email"`
	Password string `validate:"required,max=100,min=3" json:"password"`
}

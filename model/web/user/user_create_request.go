package user

type UserCreateRequest struct {
	Name     string `validate:"required,max=150,min=3" json:"name"`
	Email    string `validate:"required,max=100,min=3" json:"email"`
	Password string `validate:"required,max=100,min=3" json:"password"`
}

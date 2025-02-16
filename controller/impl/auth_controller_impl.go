package impl

import (
	"crud-dasar-go-2/controller"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/web"
	"crud-dasar-go-2/model/web/user"
	"crud-dasar-go-2/service"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	UserService service.UserService
}

// Login implements controller.AuthController.
func (authcontroller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := user.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	authcontroller.UserService.FindByEmailAndPassword(request.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK, Authenticate",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Register implements controller.AuthController.
func (authcontroller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := user.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := authcontroller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func NewAuthController(userService service.UserService) controller.AuthController {
	return &AuthControllerImpl{
		UserService: userService,
	}
}

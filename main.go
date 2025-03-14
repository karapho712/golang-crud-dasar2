package main

import (
	"crud-dasar-go-2/app"
	controller "crud-dasar-go-2/controller/impl"
	"crud-dasar-go-2/helper"
	repository "crud-dasar-go-2/repository/impl"
	service "crud-dasar-go-2/service/impl"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	kamarRepository := repository.NewKamarRepository()
	kamarService := service.NewKamarService(kamarRepository, db, validate)
	kamarController := controller.NewKamarController(kamarService)

	barangRepository := repository.NewBarangRepository()
	barangService := service.NewBarangService(barangRepository, db, validate)
	barangController := controller.NewBarangController(barangService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewAuthController(userService)

	router := app.NewRouter(kamarController, barangController, userController)

	server := http.Server{
		Addr: "localhost:3000",
		// Handler: middleware.NewAuthMiddleware(router), // temporary disable
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

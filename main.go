package main

import (
	"crud-dasar-go-2/app"
	"crud-dasar-go-2/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	server := http.Server{
		Addr: "localhost:3000",
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

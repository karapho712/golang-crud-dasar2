package exception

import (
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	if unauthorizedError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(UnauthorizedError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND ERROR",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Field(),
			})
		}

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   messages,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

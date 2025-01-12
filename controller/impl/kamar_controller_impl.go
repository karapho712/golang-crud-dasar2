package impl

import (
	"crud-dasar-go-2/controller"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/web"
	"crud-dasar-go-2/model/web/kamar"
	"crud-dasar-go-2/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type KamarControllerImpl struct {
	KamarService service.KamarService
}

// Create implements controller.KamarController.
func (kamarController *KamarControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kamarCreateRequest := kamar.KamarCreateRequest{}
	helper.ReadFromRequestBody(request, &kamarCreateRequest)

	kamarResponse := kamarController.KamarService.Create(request.Context(), kamarCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kamarResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

// Delete implements controller.KamarController.
func (kamarController *KamarControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kamarId := params.ByName("kamarId")
	id, err := strconv.Atoi(kamarId)
	helper.PanicIfError(err)

	kamarController.KamarService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll implements controller.KamarController.
func (kamarController *KamarControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kamarResponse := kamarController.KamarService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kamarResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindById implements controller.KamarController.
func (kamarController *KamarControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kamarId := params.ByName("kamarId")
	id, err := strconv.Atoi(kamarId)
	helper.PanicIfError(err)

	kamarResponse := kamarController.KamarService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kamarResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update implements controller.KamarController.
func (kamarController *KamarControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kamarUpdateRequest := kamar.KamarUpdateRequest{}
	helper.ReadFromRequestBody(request, &kamarUpdateRequest)

	kamarId := params.ByName("kamarId")
	id, err := strconv.Atoi(kamarId)
	helper.PanicIfError(err)
	kamarUpdateRequest.Id = id

	kamarResponse := kamarController.KamarService.Update(request.Context(), kamarUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kamarResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func NewKamarController(kamarService service.KamarService) controller.KamarController {
	return &KamarControllerImpl{
		KamarService: kamarService,
	}
}

package impl

import (
	"crud-dasar-go-2/controller"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/web"
	"crud-dasar-go-2/model/web/barang"
	"crud-dasar-go-2/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BarangControllerImpl struct {
	BarangService service.BarangService
}

// Create implements controller.BarangController.
func (barangController *BarangControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	barangCreateRequest := barang.BarangCreateRequest{}
	helper.ReadFromRequestBody(request, &barangCreateRequest)

	barangResponse := barangController.BarangService.Create(request.Context(), barangCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   barangResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Delete implements controller.BarangController.
func (barangController *BarangControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	barangId := params.ByName("barangId")
	id, err := strconv.Atoi(barangId)
	helper.PanicIfError(err)

	barangController.BarangService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll implements controller.BarangController.
func (barangController *BarangControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	barangResponse := barangController.BarangService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   barangResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindById implements controller.BarangController.
func (barangController *BarangControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	barangId := params.ByName("barangId")
	id, err := strconv.Atoi(barangId)
	helper.PanicIfError(err)

	barangResponse := barangController.BarangService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   barangResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update implements controller.BarangController.
func (barangController *BarangControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	barangUpdateRequest := barang.BarangUpdateRequest{}
	helper.ReadFromRequestBody(request, &barangUpdateRequest)

	barangId := params.ByName("barangId")
	id, err := strconv.Atoi(barangId)
	helper.PanicIfError(err)
	barangUpdateRequest.Id = id

	barangResponse := barangController.BarangService.Update(request.Context(), barangUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   barangResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func NewBarangController(barangService service.BarangService) controller.BarangController {
	return &BarangControllerImpl{
		BarangService: barangService,
	}
}

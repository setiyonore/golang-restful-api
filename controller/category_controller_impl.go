package controller

import (
	"net/http"
	"strconv"
	"stynore/belajar-golang-restful-api/helper"
	"stynore/belajar-golang-restful-api/model/web"
	"stynore/belajar-golang-restful-api/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryServices service.CategoryService
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)
	categoryResponse := controller.CategoryServices.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)
	categoryid := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryid)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id
	categoryResponse := controller.CategoryServices.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryid := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryid)
	helper.PanicIfError(err)

	controller.CategoryServices.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryid := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryid)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryServices.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryServices.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

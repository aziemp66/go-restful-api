package controller

import (
	"net/http"
	"strconv"

	"github.com/aziemp66/go-restful-api/helper"
	"github.com/aziemp66/go-restful-api/model/web"
	"github.com/aziemp66/go-restful-api/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (c CategoryControllerImpl) Create(
	w http.ResponseWriter,
	r *http.Request,
	p httprouter.Params,
) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, categoryCreateRequest)

	categoryResponse := c.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Value:  categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c CategoryControllerImpl) Update(
	w http.ResponseWriter,
	r *http.Request,
	p httprouter.Params,
) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, categoryUpdateRequest)

	categoryId := p.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := c.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Value:  categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c CategoryControllerImpl) Delete(
	w http.ResponseWriter,
	r *http.Request,
	p httprouter.Params,
) {
	categoryId := p.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	c.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c CategoryControllerImpl) FindById(
	w http.ResponseWriter,
	r *http.Request,
	p httprouter.Params,
) {
	categoryId := p.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := c.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Value:  categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c CategoryControllerImpl) FindAll(
	w http.ResponseWriter,
	r *http.Request,
	p httprouter.Params,
) {
	categoryResponse := c.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Value:  categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

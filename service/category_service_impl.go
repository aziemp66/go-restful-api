package service

import (
	"context"
	"database/sql"

	"github.com/aziemp66/go-restful-api/helper"
	"github.com/aziemp66/go-restful-api/model/domain"
	"github.com/aziemp66/go-restful-api/model/web"

	"github.com/aziemp66/go-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (c CategoryServiceImpl) Create(
	ctx context.Context,
	request web.CategoryCreateRequest,
) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	c.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (c CategoryServiceImpl) Update(
	ctx context.Context,
	request web.CategoryUpdateRequest,
) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = c.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (c CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	c.CategoryRepository.Delete(ctx, tx, category)
}

func (c CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (c CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := c.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryReponses(categories)
}

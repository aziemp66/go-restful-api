package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aziemp66/go-restful-api/helper"
	"github.com/aziemp66/go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (c *CategoryRepositoryImpl) Save(
	ctx context.Context,
	tx *sql.Tx,
	category domain.Category,
) domain.Category {
	SQL := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (c *CategoryRepositoryImpl) Update(
	ctx context.Context,
	tx *sql.Tx,
	category domain.Category,
) domain.Category {
	SQL := "Update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)

	helper.PanicIfError(err)

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	//TODO implement me
	SQL := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Id)

	helper.PanicIfError(err)
}

func (c *CategoryRepositoryImpl) FindById(
	ctx context.Context,
	tx *sql.Tx,
	categoryId int,
) (domain.Category, error) {
	//TODO implement me
	SQL := "select id, name from category where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, categoryId)

	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)

		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("Category Not Found")
	}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"

	rows, err := tx.QueryContext(ctx, SQL)

	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}

		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}

package main

import (
	"net/http"

	"github.com/aziemp66/go-restful-api/app"
	"github.com/aziemp66/go-restful-api/controller"
	"github.com/aziemp66/go-restful-api/helper"
	"github.com/aziemp66/go-restful-api/repository"
	"github.com/aziemp66/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	dbUser := helper.GoDotEnvVariable("DB_USER")
	dbPass := helper.GoDotEnvVariable("DB_PASS")
	dbName := helper.GoDotEnvVariable("DB_NAME")
	dbHost := helper.GoDotEnvVariable("DB_HOST")
	dbPort := helper.GoDotEnvVariable("DB_PORT")

	db := app.NewDB(dbUser, dbPass, dbName, dbHost, dbPort)
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET(("/api/categories/:id"), categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:id", categoryController.Update)
	router.DELETE(("/api/categories/:id"), categoryController.Delete)

	server := http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

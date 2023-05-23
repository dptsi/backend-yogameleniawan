package main

import (
	"net/http"

	"github.com/dptsi/backend-yogameleniawan/15-go-database-migration/app"
	"github.com/dptsi/backend-yogameleniawan/15-go-database-migration/controller"
	"github.com/dptsi/backend-yogameleniawan/15-go-database-migration/helper"
	"github.com/dptsi/backend-yogameleniawan/15-go-database-migration/middleware"
	"github.com/dptsi/backend-yogameleniawan/15-go-database-migration/repository"
	"github.com/dptsi/backend-yogameleniawan/15-go-database-migration/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

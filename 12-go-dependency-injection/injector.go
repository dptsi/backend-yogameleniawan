//go:build wireinject
// +build wireinject

package main

import (
	"github.com/dptsi/backend-yogameleniawan/12-go-dependency-injection/app"
	"github.com/dptsi/backend-yogameleniawan/12-go-dependency-injection/controller"
	"github.com/dptsi/backend-yogameleniawan/12-go-dependency-injection/middleware"
	"github.com/dptsi/backend-yogameleniawan/12-go-dependency-injection/repository"
	"github.com/dptsi/backend-yogameleniawan/12-go-dependency-injection/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

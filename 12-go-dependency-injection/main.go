package main

import (
	"net/http"

	"github.com/dptsi/backend-yogameleniawan/12-go-dependency-injection/helper"
	"github.com/dptsi/backend-yogameleniawan/12-go-dependency-injection/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

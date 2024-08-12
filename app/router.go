package app

import (
	"github.com/RoihanArrafli/golang-restful-api/controller"
	"github.com/RoihanArrafli/golang-restful-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/category", categoryController.FindAll)
	router.GET("/api/v1/category/:categoryId", categoryController.FindById)
	router.POST("/api/v1/category", categoryController.Create)
	router.PUT("/api/v1/category/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/category/:categoryId", categoryController.Delete)
	
	router.PanicHandler = exception.ErrorHandler

	return router
}
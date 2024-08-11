package router

import (
	"devbeginner-doc-api/service"
)

func InitLabsRouter() {
	router.POST("/api/lab/create", service.Labs.Create)
	router.GET("/api/lab/get", service.Labs.Get)
	router.POST("/api/lab/delete", service.Labs.Delete)
	router.POST("/api/lab/update", service.Labs.Update)
}

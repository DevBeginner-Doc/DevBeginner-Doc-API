package router

import (
	"devbeginner-doc-api/service"
)

func InitIdeRouter() {
	router.POST("/api/ide/create", service.IdeIndex.Create)
	router.GET("/api/ide/get", service.IdeIndex.Get)
	router.POST("/api/ide/delete", service.IdeIndex.Delete)
	router.POST("/api/ide/update", service.IdeIndex.Update)
}

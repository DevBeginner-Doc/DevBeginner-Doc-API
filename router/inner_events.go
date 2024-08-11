package router

import (
	"devbeginner-doc-api/service"
)

func InitInnerEventsRouter() {
	router.POST("/api/event/inner/create", service.InnerEvents.Create)
	router.GET("/api/event/inner/get", service.InnerEvents.Get)
	router.POST("/api/event/inner/delete", service.InnerEvents.Delete)
	router.POST("/api/event/inner/update", service.InnerEvents.Update)
}

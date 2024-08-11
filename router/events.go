package router

import (
	"devbeginner-doc-api/service"
)

func InitEventsRouter() {
	router.GET("/api/event/get", service.Events.Get)
}

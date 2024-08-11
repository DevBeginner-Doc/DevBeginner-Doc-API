package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	router = gin.Default()
)

func InitRouter() {
	InitLabsRouter()
	InitInnerEventsRouter()
	InitEventsRouter()
	InitIdeRouter()
	err := router.Run(fmt.Sprintf("%s:%s", viper.Get("server.addr"), viper.GetString("server.port")))
	if err != nil {
		panic(fmt.Errorf("router service fail to launch: %w", err))
	}
}

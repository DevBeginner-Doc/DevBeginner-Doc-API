package app

import (
	"devbeginner-doc-api/config"
	"devbeginner-doc-api/database"
	"devbeginner-doc-api/router"
	"fmt"
	"time"
)

func loadTimeZone() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("加载时区失败 -> ", err)
		return
	}
	time.Local = loc
}

func Run() {
	loadTimeZone()
	config.InitViper()
	database.InitDataBase()
	defer database.DB.Close()
	router.InitRouter()
}

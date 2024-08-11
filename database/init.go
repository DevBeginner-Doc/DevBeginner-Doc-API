package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var (
	DB *sqlx.DB
)

func InitDataBase() *sqlx.DB {
	sqlURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		viper.Get("mysql.username"),
		viper.GetString("mysql.password"),
		viper.Get("mysql.addr"),
		viper.GetString("mysql.port"),
		viper.Get("mysql.dbname"),
		viper.Get("mysql.charset"))
	DB = sqlx.MustConnect("mysql", sqlURL)
	fmt.Println("[DataBase] 数据库连接成功！")
	return DB
}

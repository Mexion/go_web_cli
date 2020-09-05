package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// declare a global mdb variable
var db *sqlx.DB

func Init() (err error) {
	// get the db serve name from config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	// connect db
	// or use MustConnect to panic if the connection is unsuccessful
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}

// Because db is private,
// we provide the public Close for other packages
// to close the db connection
func Close() {
	_ = db.Close()
}

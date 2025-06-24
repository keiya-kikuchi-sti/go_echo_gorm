package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	glog "github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"myapp/handler"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(glog.DEBUG)
	e.Logger.SetOutput(os.Stdout)

	db, err := run()
	if err != nil {
		e.Logger.Error(err)
		return
	}

	hd := handler.New(db)

	// "/" への GET は "hello, World!" を返す
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// "/todos/" への GET は hd.GetTodos で処理する
	e.GET("/todos/", hd.GetTodos)
	// 8080 Port で listen
	e.Logger.Info(e.Start(":8080"))
}

func run() (*gorm.DB, error) {
	dsn := "user:pass@tcp(db:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 標準出力
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}

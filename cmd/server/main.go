package main

import (
	//"database/sql"
	"ecpos/internal/config"
	"ecpos/internal/handler"
	"ecpos/internal/repository"
	"ecpos/internal/server"
	"ecpos/internal/service"
	"ecpos/pkg/http"
	"ecpos/pkg/log"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf.Env)
	addr := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	logger.Info("Server is running on " + addr)
	http.Run(logger, newApp(conf, logger), addr)
}

// newApp creates a new echo app
// and injects the dependencies
func newApp(appConfig *config.AppConfig, logger *log.Log) *echo.Echo {
	//db, err := sql.Open("mysql",
	//	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	//		appConfig.Db.UserName, appConfig.Db.Password, appConfig.Db.Host, appConfig.Db.Port, appConfig.Db.Name))

	//if err != nil {
	//	panic(err)
	//}
	newRepository := repository.NewRepository()
	userRepository := repository.NewUserRepository(newRepository)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	return server.NewServerHTTP(logger, userHandler)
}

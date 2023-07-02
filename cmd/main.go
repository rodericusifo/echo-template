package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rodericusifo/echo-template/internal/app/core"
	"github.com/rodericusifo/echo-template/internal/pkg/config"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/custom"
	"github.com/rodericusifo/echo-template/internal/pkg/util"
	"github.com/rodericusifo/echo-template/internal/pkg/util/helper"

	log "github.com/sirupsen/logrus"
)

func init() {
	config.ConfigureLog()
	config.ConfigureEnv()
	config.ConfigureDatabaseCache(constant.REDIS)
	config.ConfigureDatabaseSQL(constant.POSTGRES)
	config.ConfigureAuth()

	helper.RunDatabaseSeederSQL(constant.POSTGRES)
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = custom.CustomHTTPErrorHandler

	e.Use(
		middleware.RequestID(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           "[${time_custom}] ${remote_ip} | ${id} | ${status} | ${latency_human} | ${method} | ${path}\n",
			CustomTimeFormat: "15:04:05",
			Output:           e.Logger.Output(),
		}),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		}),
	)

	core.InitRoutes(e)

	err := e.Start(fmt.Sprintf("%s:%s", util.GetHostApp(), util.GetPortApp()))
	if err != nil {
		log.WithFields(log.Fields{
			"message": "application failed to run",
			"detail":  err,
		}).Fatal("[MAIN]")
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"pledge-bridge-backend/api/common"
	"pledge-bridge-backend/api/middlewares"
	"pledge-bridge-backend/api/routes"
	"pledge-bridge-backend/api/services"
	"pledge-bridge-backend/api/static"
	"pledge-bridge-backend/api/validate"
	"pledge-bridge-backend/config"
	"pledge-bridge-backend/db"
)

func main() {

	db.Aaaa()
	//init mysql
	db.InitMysql()

	// set next release time
	common.SetReleaseTime(config.Config.Env.ReleaseDurDev)

	// update user assets
	go services.NewUserAssetsService().UpdateUserAssets()

	//gin bind go-playground-validator
	validate.BindingValidator()

	// gin start
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	staticPath := static.GetCurrentAbPathByCaller()
	app.Static("/storage/", staticPath)
	app.Use(middlewares.Cors()) // 「 Cross domain Middleware 」
	routes.InitRoute(app)
	_ = app.Run(":" + config.Config.Env.Port)

}

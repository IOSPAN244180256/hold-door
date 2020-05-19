package main

import (
	"hold-door/config"
	"hold-door/middlewares"
	"hold-door/routers"

	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	logger := middlewares.GetDesignZapLogger()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(middlewares.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(middlewares.RecoveryWithZap(logger, true))

	//注册路由之前注册中间件
	//注册跨域中间件
	router.Use(middlewares.Cors())
	//注册gin session
	router.Use(sessions.Sessions("ginsession", middlewares.RedisSessionStore()))
	//注册auth中间件
	router.Use(middlewares.ValidataAuth())
	//注册路由
	routers.RegisterRouter(router)

	router.Run(config.GetConfig().Get("webHost.host").(string)) // listen and serve on 0.0.0.0:8080
}

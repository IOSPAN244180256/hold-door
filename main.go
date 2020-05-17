package main

import (
	"hold-door/middlewares"
	"hold-door/routers"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

func main() {
	router := gin.Default()

	//注册路由之前注册中间件

	//注册跨域中间件
	router.Use(middlewares.Cors())
	//注册gin session
	router.Use(sessions.Sessions("ginsession", middlewares.RedisSessionStore()))
	//注册auth中间件
	router.Use(middlewares.ValidataAuth())
	//注册路由
	routers.RegisterRouter(router)

	router.Run("0.0.0.0:5030") // listen and serve on 0.0.0.0:8080
}

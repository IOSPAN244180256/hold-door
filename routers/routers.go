package routers

import (
	"awesomeProject/controler/log"
	"awesomeProject/controler/sys"
	"awesomeProject/controler/trade"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	sysRouter := router.Group("/sys")
	{
		sysRouter.POST("/login", sys.Login)
	}

	logRouter := router.Group("/log")
	{
		logRouter.GET("/openhttplog", log.Queryopenhttplog)
		logRouter.GET("/operationlog", log.Queryopenhttplog)
	}

	tradeRouter := router.Group("/trade")
	{
		tradeRouter.GET("/QueryTradeDetail", trade.QueryTradeDetail)
	}
}

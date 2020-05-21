package routers

import (
	"github.com/gin-gonic/gin"
	"hold-door/controler/log"
	"hold-door/controler/sys"
	"hold-door/controler/trade"
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
		tradeRouter.GET("/querytradedetail", trade.QueryTradeDetail)
	}
}

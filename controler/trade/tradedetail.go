package trade

import (
	"context"
	"github.com/gin-gonic/gin"
	"hold-door/middlewares"
	"hold-door/models"
	GrpcTradeDetail "hold-door/protos/tradeproto"
	"hold-door/utils"
)

func QueryTradeDetail(ctx *gin.Context) {
	token := utils.Trade.MatchToken()
	if token == "" {
		logger := middlewares.GetCustomZapLogger()
		logger.Error("获取trade失败")
		ctx.JSON(200, "系统开小差了,请重试")
		return
	}

	conn, err := utils.GrpcConnWithJwt(utils.Trade, token)
	defer conn.Close()

	if err != nil {
		logger := middlewares.GetCustomZapLogger()
		logger.Error(err.Error())
		ctx.JSON(200, "系统开小差了,请重试")
		return
	}

	client := GrpcTradeDetail.NewGrpcTradeDetailServerClient(conn)
	c, cancel := context.WithTimeout(ctx.Request.Context(), models.Timeout)
	defer cancel()

	r, err2 := client.QueryTradeDetail(c, &GrpcTradeDetail.GrpcTradeDetailSel{Page: 1, Pagecount: 1000})
	if err2 != nil {
		logger := middlewares.GetCustomZapLogger()
		logger.Error(err2.Error())
		ctx.JSON(200, "系统开小差了,请重试")
		return
	}

	ctx.JSON(200, r.Rows)
}

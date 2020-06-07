package trade

import (
	"github.com/gin-gonic/gin"
	"hold-door/middlewares"
	GrpcTradeDetail "hold-door/protos/tradeproto"
	"hold-door/utils"
)

func QueryTradeDetail(ctx *gin.Context) {
	conn := utils.GrpcConnWithJwt()
	defer conn.Close()
	c := GrpcTradeDetail.NewGrpcTradeDetailServerClient(conn)
	r, err := c.QueryTradeDetail(ctx, &GrpcTradeDetail.GrpcTradeDetailSel{Page: 1, Pagecount: 1000})
	if err != nil {
		logger := middlewares.GetCustomZapLogger()
		logger.Error(err.Error())
		ctx.JSON(200, "系统开小差了,请重试")
		return
	}

	ctx.JSON(200, r.Rows)
}

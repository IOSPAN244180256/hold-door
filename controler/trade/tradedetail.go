package trade

import (
	"context"
	"github.com/gin-gonic/gin"
	"hold-door/middlewares"
	"hold-door/models"
	GrpcTradeDetail "hold-door/protos/tradeproto"
	"hold-door/utils"
	"time"
)

func QueryTradeDetail(ctx *gin.Context) {
	conn := utils.GrpcConnWithJwt()
	defer conn.Close()
	client := GrpcTradeDetail.NewGrpcTradeDetailServerClient(conn)

	c, cancel := context.WithTimeout(ctx, time.Duration(models.Timeout)*time.Second)
	defer cancel()
	r, err := client.QueryTradeDetail(c, &GrpcTradeDetail.GrpcTradeDetailSel{Page: 1, Pagecount: 1000})
	if err != nil {
		logger := middlewares.GetCustomZapLogger()
		logger.Error(err.Error())
		ctx.JSON(200, "系统开小差了,请重试")
		return
	}

	ctx.JSON(200, r.Rows)
}

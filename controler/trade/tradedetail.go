package trade

import (
	"fmt"
	"github.com/gin-gonic/gin"
	GrpcTradeDetail "hold-door/protos/tradeproto"
	"hold-door/utils"
)

func QueryTradeDetail(ctx *gin.Context) {
	conn := utils.GrpcConnWithJwt()
	defer conn.Close()
	c := GrpcTradeDetail.NewGrpcTradeDetailServerClient(conn)
	r, err := c.QueryTradeDetail(ctx, &GrpcTradeDetail.GrpcTradeDetailSel{Page: 1, Pagecount: 1})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(200, r.Rows)
}
